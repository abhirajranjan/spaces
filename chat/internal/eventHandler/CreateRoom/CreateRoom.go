package createroom

import (
	"encoding/json"
	"log"

	"github.com/abhirajranjan/spaces/chat/internal/db"
	eh "github.com/abhirajranjan/spaces/chat/internal/eventHandler"
	"github.com/abhirajranjan/spaces/chat/pkg/constants"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handle(request *kafka.Message) {
	roomreq, err := decodeRequest(request)
	if err == eh.ErrBadRequest {
		// TODO: handle bad room creation request
		log.Fatal("Bad Request")
	}
	room := db.CreateRoom(roomreq)
	if room == nil {
		// TODO: handle errcql while creating room
		log.Fatal("Server Error")
	}
	// TODO: return room created
	log.Println(room)
}

func decodeRequest(request *kafka.Message) (roomRequest *constants.RoomCreationRequest, err error) {
	if err := json.Unmarshal(request.Value, roomRequest); err != nil {
		return nil, eh.ErrBadRequest
	}
	return roomRequest, nil
}
