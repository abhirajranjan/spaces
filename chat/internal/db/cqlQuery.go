package db

var keySpace = "logs"

var roomTable = "rooms"
var roomTableStruct = " (room_id,  author_id, name, desc) "

var messagesTable = "messages"
var messagesTableStruct = " (channel_id, bucket, message_id, author_id, content) "

var userTable = "users"
var userTableStruct = " (user_id, last_snowflake_read, name) "

var namespace = keySpace + "."

var registerRoomQuery = "INSERT INTO " + namespace + roomTable + roomTableStruct + ` VALUES (%d, %d, '%s', '%s');`

var registerMessageQuery = "INSERT INTO " + namespace + messagesTable + messagesTableStruct + ` VALUES (%d, %d, %d, %d, '%s');`
var readMessageQuery = "SELECT (author_id, content, message_id) FROM " + namespace + messagesTable + ` WHERE channel_id = %d and bucket = %d;`

var InsertUserQuery = "INSERT INTO " + namespace + userTable + userTableStruct + ` VALUES (%d, %d, '%s')`
var getUserLastReadSnowFlakeQuery = "SELECT last_snowflake_read FROM " + namespace + userTable + ` WHERE user_id = %d;`
var getUserNameFromUserIDQuery = "SELECT name FROM " + namespace + userTable + ` WHERE user_id = %d;`
