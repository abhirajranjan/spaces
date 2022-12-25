package db

import (
	"context"
	"log"

	"github.com/abhirajranjan/spaces/community/config"
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/abhirajranjan/spaces/community/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
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

func find[T any](doc interface{}, opts *options.FindOptions) (*[]T, *constants.Status) {
	cur, err := collection.Find(ctx, doc, opts)
	if err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, constants.Status_ErrDb
	}
	var arr []T

	if err := cur.All(ctx, &arr); err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, constants.Status_ErrDb
	}

	return &arr, constants.Status_Ok
}

func findOne[T any](doc interface{}, opts *options.FindOneOptions) (*T, *constants.Status) {
	res := collection.FindOne(ctx, doc, opts)
	var tempHolder T
	switch res.Err() {
	case mongo.ErrNoDocuments:
		return &tempHolder, constants.Status_NoDocuments

	case nil:
		if err := res.Decode(&tempHolder); err != nil {
			logger.Logger.Sugar().Error(err)
			return nil, constants.Status_ErrDb
		}
		return &tempHolder, constants.Status_Ok

	default:
		logger.Logger.Sugar().Error(res.Err())
		return nil, constants.Status_ErrDb
	}
}

func aggregate[T any](stages []bson.D, opts *options.AggregateOptions) (*[]T, *constants.Status) {
	cur, err := collection.Aggregate(ctx, stages, opts)
	if err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, constants.Status_ErrDb
	}
	var arr []T
	if err := cur.All(ctx, &arr); err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, constants.Status_ErrDb
	}
	return &arr, constants.Status_Ok
}

func InsertOne[T any](doc interface{}, opt *options.InsertOneOptions) (*T, *constants.Status) {
	res, err := collection.InsertOne(ctx, doc, opt)
	if err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, constants.Status_ErrDb
	}
	objID := res.InsertedID.(T)
	return &objID, constants.Status_Ok
}

func FindOneAndUpdate[T any](id interface{}, update interface{}, opt *options.FindOneAndUpdateOptions) (*T, *constants.Status) {
	res := collection.FindOneAndUpdate(ctx, id, update, opt)
	if res.Err() != nil {
		logger.Logger.Sugar().Error(res.Err())
		return nil, constants.Status_ErrDb
	}
	var tempHolder T
	if err := res.Decode(&tempHolder); err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, constants.Status_ErrDb
	}
	return &tempHolder, constants.Status_Ok
}
