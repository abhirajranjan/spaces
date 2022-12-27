package db

import (
	"context"
	"log"

	"github.com/abhirajranjan/spaces/community/config"
	"github.com/abhirajranjan/spaces/community/pkg/logger"
	"github.com/abhirajranjan/spaces/community/pkg/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var community *mongo.Collection
var spaces *mongo.Collection
var ctx = context.Background()

func InitDb(database string, communityCollection string, spaceCollection string) {
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
	community = client.Database(database).Collection(communityCollection)
	spaces = client.Database(database).Collection(spaceCollection)
}

func find[T any](database string, doc interface{}, opts *options.FindOptions) (*[]T, *status.Status) {
	var cur *mongo.Cursor
	var err error
	switch database {
	case "community":
		cur, err = community.Find(ctx, doc, opts)
	case "spaces":
		cur, err = community.Find(ctx, doc, opts)
	default:
		return nil, status.ErrDb
	}
	if err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, status.ErrDb
	}
	var arr []T

	if err := cur.All(ctx, &arr); err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, status.ErrDb
	}

	return &arr, status.OkStatus
}

func findOne[T any](doc interface{}, opts *options.FindOneOptions) (*T, *status.Status) {
	res := community.FindOne(ctx, doc, opts)
	var tempHolder T
	switch res.Err() {
	case mongo.ErrNoDocuments:
		return &tempHolder, status.NoDocuments

	case nil:
		if err := res.Decode(&tempHolder); err != nil {
			logger.Logger.Sugar().Error(err)
			return nil, status.ErrDb
		}
		return &tempHolder, status.OkStatus

	default:
		logger.Logger.Sugar().Error(res.Err())
		return nil, status.ErrDb
	}
}

func aggregate[T any](stages []bson.D, opts *options.AggregateOptions) (*[]T, *status.Status) {
	cur, err := community.Aggregate(ctx, stages, opts)
	if err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, status.ErrDb
	}
	var arr []T
	if err := cur.All(ctx, &arr); err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, status.ErrDb
	}
	return &arr, status.OkStatus
}

func InsertOne[T any](doc interface{}, opt *options.InsertOneOptions) (*T, *status.Status) {
	res, err := community.InsertOne(ctx, doc, opt)
	if err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, status.ErrDb
	}
	objID := res.InsertedID.(T)
	return &objID, status.OkStatus
}

func FindOneAndUpdate[T any](id interface{}, update interface{}, opt *options.FindOneAndUpdateOptions) (*T, *status.Status) {
	res := community.FindOneAndUpdate(ctx, id, update, opt)
	if res.Err() != nil {
		logger.Logger.Sugar().Error(res.Err())
		return nil, status.ErrDb
	}
	var tempHolder T
	if err := res.Decode(&tempHolder); err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, status.ErrDb
	}
	return &tempHolder, status.OkStatus
}
