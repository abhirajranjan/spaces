package constants

import (
	"github.com/bwmarrin/snowflake"
)

type SpacePreview struct {
	Id           *snowflake.ID
	Name         string
	Tag          string
	Display_name string
	Description  string
	Banner       string
}

type Community struct {
	Id           *snowflake.ID  `bson:"_id"`
	Name         string         `bson:"name"`
	Tag          string         `bson:"tag"`
	Display_name string         `bson:"display_name"`
	Description  string         `bson:"description"`
	Banner       string         `bson:"banner"`
	Spaces       []SpacePreview `bson:"spaces"`
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
