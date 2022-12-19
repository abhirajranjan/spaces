package db

import (
	"fmt"
	"time"

	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/abhirajranjan/spaces/chat/pkg/logger"
	"github.com/abhirajranjan/spaces/chat/pkg/snowflake"
	pb "github.com/stargate/stargate-grpc-go-client/stargate/pkg/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

func registerMessage(message *constants.Message) error {
	cmd := fmt.Sprintf(registerMessageQuery,
		message.Room_id,
		make_bucket(message.Message_id),
		message.Message_id,
		message.Author_id,
		message.Content,
	)

	logger.Logger.Sugar().Debug(cmd)
	_, err := execute(nil, cmd)
	return err
}

func readMessageFromDb(message *constants.MessageRead) error {
	// TODO: gives data in reverse
	currentflake := snowflake.Generate()
	start, end := make_buckets(&currentflake, getUserLastReadSnowFlake(message.Author_id))
	var arr []*constants.MessageDocument
	for i := start; i < end; i++ {
		cmd := fmt.Sprintf(readMessageQuery, message.Room_id, i)
		pagesize := wrapperspb.Int32Value{Value: 2}
		param := pb.QueryParameters{PageSize: &pagesize, PagingState: message.PagingState}
		res, err := execute(&param, cmd)

		if err != nil {
			return err
		}

		switch a := res.Result.(type) {
		case *pb.Response_ResultSet:
			for elm := range a.ResultSet.Rows {
				row := a.ResultSet.Rows[elm].Values
				name, err := execute(nil, fmt.Sprintf(getUserNameFromUserIDQuery, row[0]))
				if err != nil {
					logger.Logger.Sugar().Error("error getting name of user %d", row[0])
					return err
				}
				doc := constants.MessageDocument{
					Author_id: snowflake.ConvertIntToSnowFlake(row[0].GetInt()),
					Content:   row[1].GetString_(),
					Time:      snowflake.ConvertIntToSnowFlake(row[2].GetInt()).Time(),
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

func getUserLastReadSnowFlake(Author_id *snowflake.ID) *snowflake.ID {
	cmd := fmt.Sprintf(getUserLastReadSnowFlakeQuery, Author_id)
	logger.Logger.Sugar().Debug(cmd)
	res, err := execute(nil, cmd)

	if err != nil {
		logger.Logger.Error(fmt.Sprintf("error getting last read snowflake for %d: %v", Author_id.Int64(), err))
		return nil
	}

	switch i := res.Result.(type) {
	case *pb.Response_ResultSet:
		// TODO: convert int64 to snowflake
		return snowflake.ConvertIntToSnowFlake(i.ResultSet.Rows[0].Values[0].GetInt())

	default:
		logger.Logger.Error(fmt.Sprintf("getting invalid type on lastreadsnowflake for %d", *Author_id))
		return nil
	}
}

func make_bucket(s *snowflake.ID) int64 {
	var timestamp int64
	if s == nil {
		timestamp = (time.Now().UnixNano() * 1000) - epoch
	} else {
		// When a Snowflake is created it contains the number of
		// seconds since the EPOCH.
		timestamp = s.Int64() >> 22
	}
	return timestamp / bucket_size
}

func make_buckets(start_id *snowflake.ID, end_id *snowflake.ID) (int64, int64) {
	return make_bucket(start_id), make_bucket(end_id) + 1
}
