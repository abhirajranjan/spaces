package db

import (
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/abhirajranjan/spaces/community/pkg/logger"
	"github.com/abhirajranjan/spaces/community/pkg/snowflake"
	"github.com/abhirajranjan/spaces/community/pkg/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewCommunity(request *constants.NewCommunityRequest) (*constants.NewCommunityResponse, *status.Status) {
	doc := bson.D{
		{Key: "name", Value: request.Name},
		{Key: "tag", Value: request.Tag},
		{Key: "description", Value: request.Description},
		{Key: "banner", Value: request.Banner},
	}
	object_id, status := InsertOne[snowflake.ID](NSCommunity, doc, nil)
	response := constants.NewCommunityResponse{
		Id:           object_id,
		Name:         request.Name,
		Tag:          request.Tag,
		Display_name: request.Display_name,
		Description:  request.Description,
		Banner:       request.Banner,
	}
	return &response, status
}

func UpdateCommunity(request *constants.UpdatedCommunityRequest) (*constants.UpdatedCommunityResponse, *status.Status) {
	tempHolder := bson.M{}
	if request.Banner != "" {
		tempHolder["banner"] = request.Banner
	}
	if request.Description != "" {
		tempHolder["description"] = request.Description
	}
	if request.Display_name != "" {
		tempHolder["display_name"] = request.Display_name
	}
	if request.Name != "" {
		tempHolder["name"] = request.Name
	}
	if request.Tag != "" {
		tempHolder["tag"] = request.Tag
	}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	return FindOneAndUpdate[constants.UpdatedCommunityResponse](NSCommunity, bson.M{"id": request.Id}, tempHolder, &opt)
}

func UpdateSpaces(request *constants.UpdateSpaceRequest) (*constants.UpdateSpaceResponse, *status.Status) {
	switch request.Type {
	case constants.Global_space:
		return handleGlobalSpace(request)

	case constants.Text_space:
		return handleTextSpace(request)

	case constants.Video_space:
		return handleVideoSpace(request)

	case constants.Image_space:
		return handleImageSpace(request)

	case constants.Embedding_space:
		return handleEmbeddingSpace(request)

	default:
		logger.Logger.Sugar().Error("unknown space type provided", request)
		return nil, status.GenerateBadRequest("unknown space type provided")
	}
}

func handleGlobalSpace(request *constants.UpdateSpaceRequest) (*constants.UpdateSpaceResponse, *status.Status) {
	doc := bson.M{}

	if request.Id == 0 {
		return nil, status.GenerateBadRequest("id not specified")
	}

	if request.Name != "" && request.Tag != "" {
		doc["name"] = request.Name
		doc["tag"] = request.Tag
	}
	if request.Display_name != "" {
		doc["display_name"] = request.Display_name
	}
	if request.Description != "" {
		doc["description"] = request.Description
	}
	if request.Banner != "" {
		doc["banner"] = request.Banner
	}

	rd := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &rd,
	}

	return FindOneAndUpdate[constants.UpdateSpaceResponse](NSSpace, bson.M{"_id": request.Id}, doc, &opt)
}

func handleTextSpace(request *constants.UpdateSpaceRequest) (*constants.UpdateSpaceResponse, *status.Status) {
	doc := bson.M{}

	if request.Id == 0 {
		return nil, status.GenerateBadRequest("id not specified")
	}
	if request.Name != "" && request.Tag != "" {
		doc["name"] = request.Name
		doc["tag"] = request.Tag
	}

	if request.Markdown != "" {
		doc["markdown"] = request.Markdown
	}

	rd := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &rd,
	}

	return FindOneAndUpdate[constants.UpdateSpaceResponse](NSSpace, bson.M{"_id": request.Id}, doc, &opt)
}

func handleVideoSpace(request *constants.UpdateSpaceRequest) (*constants.UpdateSpaceResponse, *status.Status) {
	doc := bson.M{}

	if request.Id == 0 {
		return nil, status.GenerateBadRequest("id not specified")
	}

	if request.Name != "" && request.Tag != "" {
		doc["name"] = request.Name
		doc["tag"] = request.Tag
	}

	if request.Src != "" {
		doc["src"] = request.Src
	}

	if request.Format != "" {
		doc["format"] = request.Format
	}

	rd := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &rd,
	}

	return FindOneAndUpdate[constants.UpdateSpaceResponse](NSSpace, bson.M{"_id": request.Id}, doc, &opt)
}

func handleImageSpace(request *constants.UpdateSpaceRequest) (*constants.UpdateSpaceResponse, *status.Status) {
	doc := bson.M{}

	if request.Id == 0 {
		return nil, status.GenerateBadRequest("id not specified")
	}

	if request.Name != "" && request.Tag != "" {
		doc["name"] = request.Name
		doc["tag"] = request.Tag
	}

	if request.Src != "" {
		doc["src"] = request.Src
	}

	if request.Format != "" {
		doc["format"] = request.Format
	}

	rd := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &rd,
	}

	return FindOneAndUpdate[constants.UpdateSpaceResponse](NSSpace, bson.M{"_id": request.Id}, doc, &opt)
}

func handleEmbeddingSpace(request *constants.UpdateSpaceRequest) (*constants.UpdateSpaceResponse, *status.Status) {
	doc := bson.M{}

	if request.Id == 0 {
		return nil, status.GenerateBadRequest("id not specified")
	}

	if request.Name != "" && request.Tag != "" {
		doc["name"] = request.Name
		doc["tag"] = request.Tag
	}

	if request.Src != "" {
		doc["src"] = request.Src
	}

	rd := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &rd,
	}

	return FindOneAndUpdate[constants.UpdateSpaceResponse](NSSpace, bson.M{"_id": request.Id}, doc, &opt)
}
