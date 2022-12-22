package constants

import (
	"github.com/bwmarrin/snowflake"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Message struct {
	Room_id    *snowflake.ID `json:"room_id"`
	Message_id *snowflake.ID `json:"message_id"`
	Author_id  *snowflake.ID `json:"author_id"`
	Content    string        `json:"content"`
}

type MessageRead struct {
	Room_id     *snowflake.ID          `json:"room_id"`
	Author_id   *snowflake.ID          `json:"author_id"`
	PageSize    *wrapperspb.Int32Value `json:"page_size"`
	PagingState *wrapperspb.BytesValue `json:"paging_state"`
	Content     []*MessageDocument
}

type MessageDocument struct {
	Author_id *snowflake.ID
	Name      string
	Content   string
	Time      int64
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

// passive listeners
type NewUserCreated struct {
	User_id *snowflake.ID `json:"user_id"`
	Name    string        `json:"name"`
}
