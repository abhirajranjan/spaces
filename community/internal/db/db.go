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

const (
	NSCommunity = "community"
	NSSpace     = "spaces"
)

var (
	community *mongo.Collection
	spaces    *mongo.Collection
	ctx       = context.Background()
)

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
	case NSCommunity:
		cur, err = community.Find(ctx, doc, opts)
	case NSSpace:
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

	return &arr, status.Ok
}

func findOne[T any](database string, doc interface{}, opts *options.FindOneOptions) (*T, *status.Status) {
	var res *mongo.SingleResult

	switch database {
	case NSCommunity:
		res = community.FindOne(ctx, doc, opts)
	case NSSpace:
		res = spaces.FindOne(ctx, doc, opts)
	default:
		return nil, status.ErrDb
	}

	var tempHolder T
	switch res.Err() {
	case mongo.ErrNoDocuments:
		return &tempHolder, status.NoDataFound

	case nil:
		if err := res.Decode(&tempHolder); err != nil {
			logger.Logger.Sugar().Error(err)
			return nil, status.ErrDb
		}
		return &tempHolder, status.Ok

	default:
		logger.Logger.Sugar().Error(res.Err())
		return nil, status.ErrDb
	}
}

func aggregate[T any](database string, stages []bson.D, opts *options.AggregateOptions) (*[]T, *status.Status) {
	var cur *mongo.Cursor
	var err error

	switch database {
	case NSCommunity:
		cur, err = community.Aggregate(ctx, stages, opts)
	case NSSpace:
		cur, err = spaces.Aggregate(ctx, stages, opts)
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
	return &arr, status.Ok
}

func InsertOne[T any](database string, doc interface{}, opt *options.InsertOneOptions) (*T, *status.Status) {
	var res *mongo.InsertOneResult
	var err error

	switch database {
	case NSCommunity:
		res, err = community.InsertOne(ctx, doc, opt)
	case NSSpace:
		res, err = community.InsertOne(ctx, doc, opt)
	default:
		logger.Logger.Error("unknown database specified")
		return nil, status.ErrDb
	}

	if err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, status.ErrDb
	}
	objID := res.InsertedID.(T)
	return &objID, status.Ok
}

func FindOneAndUpdate[T any](database string, id interface{}, update interface{}, opt *options.FindOneAndUpdateOptions) (*T, *status.Status) {
	var res *mongo.SingleResult

	switch database {
	case NSCommunity:
		res = community.FindOneAndUpdate(ctx, id, update, opt)
	case NSSpace:
		res = spaces.FindOneAndUpdate(ctx, id, update, opt)
	}

	if res.Err() != nil {
		logger.Logger.Sugar().Error(res.Err())
		return nil, status.ErrDb
	}
	var tempHolder T
	if err := res.Decode(&tempHolder); err != nil {
		logger.Logger.Sugar().Error(err)
		return nil, status.ErrDb
	}
	return &tempHolder, status.Ok
}

func getSpaces[T any](id interface{}, opt *options.FindOneOptions) (*T, *status.Status) {
	return findOne[T](NSSpace, id, opt)
}
