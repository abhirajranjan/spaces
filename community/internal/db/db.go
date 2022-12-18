package db

import (
	"context"
	"log"

	"github.com/abhirajranjan/spaces/community/config"
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/abhirajranjan/spaces/community/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.Background()

func InitDb(database string, coll string) {
	clientOptions := options.Client().ApplyURI(config.DbURI())
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Logger.Sugar().Warn(err)
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Logger.Sugar().Warn(err)
		log.Fatal(err)
	}
	collection = client.Database(database).Collection(coll)
}

func CommunityRequestToBson(json constants.Community) (d bson.D) {
	if json.Id == 0 {
		d = append(d, primitive.E{Key: "id", Value: json.Id})
	}
	if json.Name == "" {
		d = append(d, primitive.E{Key: "name", Value: json.Name})
	}
	if json.Tag == "" {
		d = append(d, primitive.E{Key: "tag", Value: json.Tag})
	}
	return d
}
