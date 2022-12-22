package db

import (
	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/abhirajranjan/spaces/chat/pkg/logger"
	"github.com/abhirajranjan/spaces/chat/pkg/snowflake"
)

func CreateRoom(request *constants.RoomCreationRequest) (room *constants.Room) {
	id := snowflake.Generate()
	room.Room_id = &id
	room.Name = request.Name
	room.Desc = request.Desc
	room.Author_id = request.Author_id
	if err := registerRoom(room); err != nil {
		return nil
	}
	return room
}

func registerRoom(room *constants.Room) error {
	cmd := RegisterRoomQuery(room.Room_id.Int64(), room.Author_id.Int64(), room.Name, room.Desc)
	logger.Logger.Sugar().Debugln(cmd)
	if _, err := execute(nil, cmd); err == ErrCql {
		return err
	}
	return nil
}
