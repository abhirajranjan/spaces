package db

import (
	"github.com/abhirajranjan/spaces/community/pkg/constants"
	"github.com/abhirajranjan/spaces/community/pkg/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SearchCommunity(request *constants.SearchCommunityRequest) (*[]constants.SearchCommunityResponse, *status.Status) {
	stages := []bson.D{
		// match stage
		{{
			Key: "$and",
			Value: bson.A{
				bson.D{{
					Key: "name",
					Value: bson.D{{
						Key:   "$regex",
						Value: "^.*" + request.Name + ".*$",
					}},
				}},

				bson.D{{
					Key: "tag",
					Value: bson.D{{
						Key:   "$regex",
						Value: "^.*" + request.Tag + ".*$",
					}},
				}},
			},
		}},
		// limit stage
		{{
			Key:   "$limit",
			Value: request.Pagesize,
		}},
	}

	ps := int32(request.Pagesize)
	opts := options.AggregateOptions{
		BatchSize: &ps,
	}
	return aggregate[constants.SearchCommunityResponse](stages, &opts)
}

func GetCommunity(request *constants.GetCommunityRequest) (*constants.GetCommunityResponse, *status.Status) {
	doc := bson.D{{
		Key: "$and",
		Value: bson.A{
			bson.D{{
				Key:   "name",
				Value: request.Name,
			}},

			bson.D{{
				Key:   "tag",
				Value: request.Tag,
			}},
		},
	}}

	return findOne[constants.GetCommunityResponse](doc, nil)
}
