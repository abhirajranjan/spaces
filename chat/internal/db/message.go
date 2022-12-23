package db

import (
	"fmt"

	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/abhirajranjan/spaces/chat/pkg/logger"
	"github.com/abhirajranjan/spaces/chat/pkg/snowflake"
	pb "github.com/stargate/stargate-grpc-go-client/stargate/pkg/proto"
)

var bucket_size int64 = 1000 * 60 * 60 * 24 * 10

func CreateMessage(messagereq *constants.MessageRequest) (message *constants.Message, status *constants.Status) {
	id := snowflake.Generate()

	message.Author_id = messagereq.Author_id
	message.Room_id = messagereq.Room_id
	message.Content = messagereq.Content
	message.Message_id = &id

	status = registerMessage(message)
	return message, status
}

func ReadMessage(messagereq *constants.MessageReadRequest) (message *constants.MessageRead, status *constants.Status) {
	message.Author_id = messagereq.Author_id
	message.PageSize = messagereq.PageSize
	message.PagingState = messagereq.PagingState
	message.Room_id = messagereq.Room_id

	status = readMessageFromDb(message)
	return message, status
}

func DeleteChat(messagereq *constants.MessageDeleteRequest) (status *constants.Status) {
	cmd := DeleteMessageQuery(
		messagereq.Message_id.Int64(),
		messagereq.Room_id.Int64(),
		messagereq.Bucket,
	)
	logger.Logger.Sugar().Debug(cmd)
	_, status = execute(nil, cmd)
	return status
}

func DeleteRoom(messagereq *constants.DeleteRoomRequest) (status *constants.Status) {
	cmd := []string{
		DeleteRoomQuery(messagereq.Room_id.Int64()),
		DeleteLastReadsOnRoomDeleteQuery(messagereq.Room_id.Int64()),
	}
	start, stop := make_buckets(messagereq.Room_id, nil)
	for i := start; start < stop; start++ {
		cmd = append(cmd, DeleteMessageOnRoomDeleteQuery(messagereq.Room_id.Int64(), i))
	}
	logger.Logger.Sugar().Debug(cmd)
	_, status = executeMultiple(nil, cmd)
	return status
}

func registerMessage(message *constants.Message) *constants.Status {
	cmd := RegisterMessageQuery(
		message.Room_id.Int64(),
		make_bucket(message.Message_id),
		message.Message_id.Int64(),
		message.Author_id.Int64(),
		message.Content,
	)

	logger.Logger.Sugar().Debug(cmd)
	_, status := execute(nil, cmd)
	return status
}

// TODO: gives data in reverse
func readMessageFromDb(message *constants.MessageRead) *constants.Status {
	currentflake := snowflake.Generate()
	lastmessageread := getUserLastReadSnowFlake(message.Author_id, message.Room_id)
	if lastmessageread == nil {
		lastmessageread = message.Room_id
	}
	start, end := make_buckets(lastmessageread, &currentflake)

	var arr []*constants.MessageDocument
	for bucket := start; bucket < end; bucket++ {
		cmd := ReadMessageQuery(message.Room_id.Int64(), bucket)
		param := pb.QueryParameters{PageSize: message.PageSize, PagingState: message.PagingState}
		res, status := execute(&param, cmd)

		if status.Value != constants.Ok {
			logger.Logger.Sugar().Error("cql error: cannot read message", status)
			return constants.Status_ErrCql
		}

		switch a := res.Result.(type) {
		case *pb.Response_ResultSet:
			for elm := range a.ResultSet.Rows {
				row := a.ResultSet.Rows[elm].Values
				name, status := execute(nil, cmd)
				cmd = GetUserNameFromUserIDQuery(message.Author_id.Int64())

				if len(name.GetResultSet().Rows) == 0 || len(name.GetResultSet().Rows[0].Values) == 0 || status.Value != constants.Ok {
					logger.Logger.Sugar().Error("error getting name of user %d", row[0])
					return constants.Status_ErrCql
				}
				doc := constants.MessageDocument{
					Name:    name.GetResultSet().Rows[0].Values[0].GetString_(),
					Content: row[1].GetString_(),
					Time:    snowflake.ParseInt64(row[2].GetInt()).Time(),
					Bucket:  bucket,
				}
				arr = append(arr, &doc)
			}
		default:
			logger.Logger.Sugar().Errorf("undefined type while read query: %s", cmd)
			return constants.Status_ErrCql
		}
	}
	message.Content = arr
	return constants.Status_Ok
}

func getUserLastReadSnowFlake(Author_id *snowflake.ID, Room_id *snowflake.ID) *snowflake.ID {
	cmd := GetUserLastReadSnowFlakeQuery(Author_id.Int64(), Room_id.Int64())
	logger.Logger.Sugar().Debug(cmd)
	res, err := execute(nil, cmd)

	if err != nil {
		logger.Logger.Error(fmt.Sprintf("error getting last read snowflake for %d: %v", Author_id.Int64(), err))
		return nil
	}

	switch i := res.Result.(type) {
	case *pb.Response_ResultSet:
		if len(i.ResultSet.Rows) == 0 || len(i.ResultSet.Rows[0].Values) == 0 {
			return nil
		}
		return snowflake.ParseInt64(i.ResultSet.Rows[0].Values[0].GetInt())

	default:
		logger.Logger.Error(fmt.Sprintf("getting invalid type on lastreadsnowflake for %d", *Author_id))
		return nil
	}
}

func make_bucket(s *snowflake.ID) int64 {
	var timestamp int64
	if s == nil {
		timestamp = snowflake.Generate().Time()
	} else {
		// When a Snowflake is created it contains the number of
		// seconds since the EPOCH.
		timestamp = s.Time()
	}
	return timestamp / bucket_size
}

func make_buckets(start_id *snowflake.ID, end_id *snowflake.ID) (start int64, end int64) {
	return make_bucket(start_id), make_bucket(end_id) + 1
}
