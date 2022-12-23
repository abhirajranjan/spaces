package db

import (
	"fmt"
	"time"

	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/abhirajranjan/spaces/chat/pkg/logger"
	"github.com/abhirajranjan/spaces/chat/pkg/snowflake"
	pb "github.com/stargate/stargate-grpc-go-client/stargate/pkg/proto"
)

var epoch int64 = 1420070400000
var bucket_size int64 = 1000 * 60 * 60 * 24 * 10

func CreateMessage(messagereq *constants.MessageRequest) (message *constants.Message) {
	id := snowflake.Generate()

	message.Author_id = messagereq.Author_id
	message.Room_id = messagereq.Room_id
	message.Content = messagereq.Content
	message.Message_id = &id

	if err := registerMessage(message); err != nil {
		logger.Logger.Sugar().Error("error registering message", err, *message)
		return nil
	}
	return message
}

func ReadMessage(messagereq *constants.MessageReadRequest) (message *constants.MessageRead) {
	message.Author_id = messagereq.Author_id
	message.PageSize = messagereq.PageSize
	message.PagingState = messagereq.PagingState
	message.Room_id = messagereq.Room_id

	readMessageFromDb(message)

	return message
}

func DeleteChat(messagereq *constants.MessageDeleteRequest) (message *constants.MessageDelete) {
	message.Author_id = messagereq.Author_id
	message.Bucket = messagereq.Bucket
	message.Message_id = messagereq.Message_id
	message.Room_id = messagereq.Room_id

	if err := deleteMessageFromDb(message); err != nil {
		logger.Logger.Sugar().Error("error deleting message", err, *message)
		message.Status = 300
	} else {
		message.Status = 200
	}
	return message
}

func DeleteRoom(messagereq *constants.DeleteRoomRequest) (message *constants.DeleteRoom) {
	message.Author_id = messagereq.Author_id
	message.Room_id = messagereq.Author_id

	if err := deleteRoomFromDb(message); err != nil {
		message.Status = 200
	} else {
		message.Status = 200
	}
	return message
}

func registerMessage(message *constants.Message) error {
	cmd := RegisterMessageQuery(
		message.Room_id.Int64(),
		make_bucket(message.Message_id),
		message.Message_id.Int64(),
		message.Author_id.Int64(),
		message.Content,
	)

	logger.Logger.Sugar().Debug(cmd)
	_, err := execute(nil, cmd)
	return err
}

// TODO: gives data in reverse
func readMessageFromDb(message *constants.MessageRead) error {
	currentflake := snowflake.Generate()
	lastmessageread := getUserLastReadSnowFlake(message.Author_id, message.Room_id)
	if lastmessageread == nil {
		lastmessageread = message.Room_id
	}
	start, end := make_buckets(lastmessageread, &currentflake)

	var arr []*constants.MessageDocument
	for i := start; i < end; i++ {
		cmd := ReadMessageQuery(message.Room_id.Int64(), i)
		param := pb.QueryParameters{PageSize: message.PageSize, PagingState: message.PagingState}
		res, err := execute(&param, cmd)

		if err != nil {
			return err
		}

		switch a := res.Result.(type) {
		case *pb.Response_ResultSet:
			for elm := range a.ResultSet.Rows {
				row := a.ResultSet.Rows[elm].Values
				name, err := execute(nil, cmd)
				cmd = GetUserNameFromUserIDQuery(message.Author_id.Int64())

				if len(name.GetResultSet().Rows) == 0 || len(name.GetResultSet().Rows[0].Values) == 0 || err != nil {
					logger.Logger.Sugar().Error("error getting name of user %d", row[0])
					return err
				}
				doc := constants.MessageDocument{
					Author_id: snowflake.ParseInt64(row[0].GetInt()),
					Content:   row[1].GetString_(),
					Time:      snowflake.ParseInt64(row[2].GetInt()).Time(),
					Name:      name.GetResultSet().Rows[0].Values[0].GetString_(),
				}
				arr = append(arr, &doc)
			}
		default:
			logger.Logger.Sugar().Errorf("undefined type while read query: %s", cmd)
			return ErrCql
		}
	}
	message.Content = arr
	return nil
}

func deleteMessageFromDb(message *constants.MessageDelete) error {
	cmd := DeleteMessageQuery(
		message.Message_id.Int64(),
		message.Room_id.Int64(),
		message.Bucket,
	)
	logger.Logger.Sugar().Debug(cmd)
	_, err := execute(nil, cmd)
	return err
}

func deleteRoomFromDb(message *constants.DeleteRoom) error {
	cmd := []string{
		DeleteRoomQuery(message.Room_id.Int64()),
		DeleteLastReadsOnRoomDeleteQuery(message.Room_id.Int64()),
	}
	start, stop := make_buckets(message.Room_id, nil)
	for i := start; start < stop; start++ {
		cmd = append(cmd, DeleteMessageOnRoomDeleteQuery(message.Room_id.Int64(), i))
	}
	_, err := executeMultiple(nil, cmd)
	return err
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
		timestamp = (time.Now().Unix() * 1000) - epoch
	} else {
		// When a Snowflake is created it contains the number of
		// seconds since the EPOCH.
		timestamp = s.Int64() >> 22
	}
	return timestamp / bucket_size
}

func make_buckets(start_id *snowflake.ID, end_id *snowflake.ID) (start int64, end int64) {
	return make_bucket(start_id), make_bucket(end_id) + 1
}
