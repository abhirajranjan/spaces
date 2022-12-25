package constants

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status struct {
	Value int
	Err   *error
}

func (d *Status) Error() string {
	return fmt.Sprint(d.Err)
}

func NewStatus(value int, v any) *Status {
	var err error
	switch cus := v.(type) {
	case string:
		err = fmt.Errorf(cus)
	case error:
		err = cus
	default:
		return nil
	}
	return &Status{
		Value: value,
		Err:   &err,
	}
}

type SpacePreview struct {
	Id           *primitive.ObjectID
	Name         string
	Tag          string
	Display_name string
	Description  string
	Banner       string
}

type Community struct {
	Id           *primitive.ObjectID `bson:"_id"`
	Name         string              `bson:"name"`
	Tag          string              `bson:"tag"`
	Display_name string              `bson:"display_name"`
	Description  string              `bson:"description"`
	Banner       string              `bson:"banner"`
	Spaces       []SpacePreview      `bson:"spaces"`
}

// requests

type SearchCommunityRequest struct {
	Name     string `json:"name"`
	Tag      string `json:"tag"`
	Pagesize int    `json:"pagesize"`
}
type GetCommunityRequest struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type NewCommunityRequest struct {
	Name         string `json:"namespace_name"`
	Tag          string `json:"namespace_tag"`
	Display_name string `json:"name"`
	Description  string `json:"description"`
	Banner       string `json:"banner"`
}

type UpdatedCommunityRequest Community

// response

type GetCommunityResponse Community
type SearchCommunityResponse Community
type NewCommunityResponse Community
type UpdatedCommunityResponse Community
