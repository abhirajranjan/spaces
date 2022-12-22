package constants

// lists the topic event utilised in the service
var Self = "chat"

// get
var Event_ReadChat = "chat.ReadChat"

// post
var Event_NewChat = "chat.NewChat"
var Event_DeleteChat = "chat.DeleteChat"
var Event_CreateRoom = "chat.CreateRoom"
var Event_DeleteRoom = "chat.DeleteRoom"
var Event_NewUserCreated = "user.NewUserCreated"

var ConsumeGroupID = "chat"
var Log = "logging"

// enviroment variables
// * kafka config *
var ENV_producerConfig = "KAFKA_PRODUCER_CONFIG"
var ENV_consumerConfig = "KAFKA_CONSUMER_CONFIG"

// * AstraDB config *
var ENV_astraClusterId = "ASTRA_CLUSTER_ID"
var ENV_astraRegion = "ASTRA_REGION"
var ENV_BearerToken = "ASTRA_BEARER_TOKEN"

var DB_URI = "%s-%s.apps.astra.datastax.com:443"
