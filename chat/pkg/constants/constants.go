package constants

// lists the topic event utilised in the service

var (
	Self = "chat"

	KafkaListenOn = []string{
		Self,
		"user",
	}
	// get
	Event_ReadChat = "chat.ReadChat"

	// post
	Event_NewChat        = "chat.NewChat"
	Event_DeleteChat     = "chat.DeleteChat"
	Event_CreateRoom     = "chat.CreateRoom"
	Event_DeleteRoom     = "chat.DeleteRoom"
	Event_NewUserCreated = "user.NewUserCreated"

	ConsumeGroupID = "chat"
	Log            = "logging"

	// enviroment variables
	// * kafka config *
	ENV_producerConfig = "KAFKA_PRODUCER_CONFIG"
	ENV_consumerConfig = "KAFKA_CONSUMER_CONFIG"

	// * AstraDB config *
	ENV_astraClusterId = "ASTRA_CLUSTER_ID"
	ENV_astraRegion    = "ASTRA_REGION"
	ENV_BearerToken    = "ASTRA_BEARER_TOKEN"
	DB_URI             = "%s-%s.apps.astra.datastax.com:443"
)
