package db

import (
	"github.com/abhirajranjan/spaces/community/pkg/constants"
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
	object_id, status := InsertOne[snowflake.ID](doc, nil)
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
	return FindOneAndUpdate[constants.UpdatedCommunityResponse](bson.M{"id": request.Id}, tempHolder, &opt)
}
