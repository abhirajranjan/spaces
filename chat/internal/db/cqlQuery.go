package db

import (
	"fmt"
)

var keySpace = "logs"

var roomTable = "rooms"
var roomTableStruct = " (room_id,  author_id, name, description) "

var messagesTable = "messages"
var messagesTableStruct = " (channel_id, bucket, message_id, author_id, content) "

var userTable = "users"
var userTableStruct = " (user_id, name) "

var userLastReadTable = "userlastread"
var userLastReadTableStruct = " (user_id, channel_id, last_snowflake_read) "

var namespace = keySpace + "."

var registerRoomQuery = "INSERT INTO " + namespace + roomTable + roomTableStruct + ` VALUES (%d, %d, '%s', '%s');`

var registerMessageQuery = "INSERT INTO " + namespace + messagesTable + messagesTableStruct + ` VALUES (%d, %d, %d, %d, '%s');`
var readMessageQuery = "SELECT author_id, content, message_id FROM " + namespace + messagesTable + ` WHERE channel_id = %d and bucket = %d;`
var deleteMessageQuery = " DELETE FROM " + namespace + messagesTable + ` WHERE message_id = %d and channel_id = %d and bucket = %d;`

var insertUserQuery = "INSERT INTO " + namespace + userTable + userTableStruct + ` VALUES (%d, '%s');`
var getUserNameFromUserIDQuery = "SELECT name FROM " + namespace + userTable + ` WHERE user_id = %d;`

var updateUserLastReadQuery = "INSERT INTO " + namespace + userLastReadTable + userLastReadTableStruct + ` VALUES (%d, %d, %d);`
var getUserLastReadSnowFlakeQuery = "SELECT last_snowflake_read FROM " + namespace + userLastReadTable + ` WHERE user_id = %d; and channel_id = %d;`

func RegisterRoomQuery(room_id int64, author_id int64, name string, description string) string {
	return fmt.Sprintf(registerRoomQuery, room_id, author_id, name, description)
}

func RegisterMessageQuery(channel_id int64, bucket int64, message_id int64, author_id int64, content string) string {
	return fmt.Sprintf(registerMessageQuery, channel_id, bucket, message_id, author_id, content)
}

func ReadMessageQuery(channel_id int64, bucket int64) string {
	return fmt.Sprintf(readMessageQuery, channel_id, bucket)
}

func DeleteMessageQuery(message_id int64, channel_id int64, bucket int64) string {
	return fmt.Sprintf(deleteMessageQuery, message_id, channel_id, bucket)
}

func InsertUserQuery(user_id int64, name string) string {
	return fmt.Sprintf(insertUserQuery, user_id, name)
}

func GetUserNameFromUserIDQuery(user_id int64) string {
	return fmt.Sprintf(getUserNameFromUserIDQuery, user_id)
}

func UpdateUserLastReadQuery(user_id int64, channel_id int64, last_snowflake_read int64) string {
	return fmt.Sprintf(updateUserLastReadQuery, user_id, channel_id, last_snowflake_read)
}

func GetUserLastReadSnowFlakeQuery(user_id int64, channel_id int64) string {
	return fmt.Sprintf(getUserLastReadSnowFlakeQuery, user_id, channel_id)
}
