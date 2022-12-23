package db

import (
	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/abhirajranjan/spaces/chat/pkg/logger"
	"github.com/abhirajranjan/spaces/chat/pkg/snowflake"
)

func CreateRoom(request *constants.RoomCreationRequest) (room *constants.Room, status *constants.Status) {
	id := snowflake.Generate()
	room.Room_id = &id
	room.Name = request.Name
	room.Desc = request.Desc
	room.Author_id = request.Author_id

	status = registerRoom(room)
	return room, status
}

func registerRoom(room *constants.Room) *constants.Status {
	cmd := RegisterRoomQuery(room.Room_id.Int64(), room.Author_id.Int64(), room.Name, room.Desc)
	logger.Logger.Sugar().Debugln(cmd)
	// execute only returns status codes that can be handled by event handler
	_, status := execute(nil, cmd)
	return status
}
