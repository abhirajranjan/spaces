package constants

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Status struct {
	Value int
	Err   *error
}

func (d *Status) Error() string {
	return fmt.Sprint(d.Err)
}

func NewStatus(value int, v any) *Status {
	var err error
	switch cus := v.(type) {
	case string:
		err = fmt.Errorf(cus)
	case error:
		err = cus
	default:
		return nil
	}
	return &Status{
		Value: value,
		Err:   &err,
	}
}

// response

type Message struct {
	Room_id    *snowflake.ID `json:"room_id"`
	Message_id *snowflake.ID `json:"message_id"`
	Author_id  *snowflake.ID `json:"author_id"`
	Content    string        `json:"content"`
	Bucket     int64         `json:"bucket"`
}

type MessageRead struct {
	Room_id     *snowflake.ID          `json:"room_id"`
	Author_id   *snowflake.ID          `json:"author_id"`
	PageSize    *wrapperspb.Int32Value `json:"page_size"`
	PagingState *wrapperspb.BytesValue `json:"paging_state"`
	Content     []*MessageDocument
}

// TODO: BUG: dont give author id else one can replicate them
type MessageDocument struct {
	Author_id *snowflake.ID `json:"author_id"`
	Name      string        `json:"name"`
	Content   string        `json:"content"`
	Time      int64         `json:"time"`
	Bucket    int64         `json:"bucket"`
}

type User struct {
	User_id *snowflake.ID `json:"user_id"`
}

type Room struct {
	Room_id   *snowflake.ID `json:"room_id"`
	Name      string        `json:"name"`
	Desc      string        `json:"desc"`
	Author_id *snowflake.ID `json:"author_id"`
}

// requests

type RoomCreationRequest struct {
	Name      string        `json:"name"`
	Desc      string        `json:"description"`
	Author_id *snowflake.ID `json:"author_id"`
}

type MessageRequest struct {
	Room_id   *snowflake.ID `json:"room_id"`
	Author_id *snowflake.ID `json:"author_id"`
	Content   string        `json:"content"`
}

type MessageReadRequest struct {
	Room_id     *snowflake.ID          `json:"room_id"`
	Author_id   *snowflake.ID          `json:"author_id"`
	PageSize    *wrapperspb.Int32Value `json:"page_size"`
	PagingState *wrapperspb.BytesValue `json:"paging_state"`
}

type MessageDeleteRequest struct {
	Room_id    *snowflake.ID `json:"room_id"`
	Author_id  *snowflake.ID `json:"author_id"`
	Message_id *snowflake.ID `json:"message_id"`
	Bucket     int64         `json:"bucket"`
}

type DeleteRoomRequest struct {
	Room_id   *snowflake.ID `json:"room_id"`
	Author_id *snowflake.ID `json:"author_id"`
}

// passive listeners

type NewUserCreated struct {
	User_id *snowflake.ID `json:"user_id"`
	Name    string        `json:"name"`
}
