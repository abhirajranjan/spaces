package constants

// lists the topic event utilised in the service
var Self string = "community"

// get events
var GetCommunity string = "community.GetCommunity"
var SearchCommunity string = "community.SearchCommunity"

// post events
var NewCommunity string = "community.NewCommunity"
var UpdateCommunity string = "community.UpdateCommunity"

var ConsumeGroupID string = "community"
var Log string = "logging"

// enviroment variables
var ENV_databaseURI = "db_uri"
var ENV_producerConfig = "kafka_producer_config"
var ENV_consumerConfig = "kafka_consumer_config"
