package constants

import (
	"github.com/bwmarrin/snowflake"
)

const (
	Global_space = iota
	Text_space
	Video_space
	Image_space
	Embedding_space
)

// spaces
type (
	Space struct {
		Id           int64
		Type         int
		Name         string
		Tag          string
		Markdown     string
		Src          string
		Format       string
		Display_name string
		Description  string
		Banner       string
	}

	SpacePreview struct {
		Id           int64  `json:"_id"`
		Type         int    // 0
		Name         string `json:"name"`
		Tag          string `json:"tag"`
		Display_name string `json:"display_name"`
		Description  string `json:"description"`
		Banner       string `json:"banner"`
	}

	TextSpace struct {
		Id       int64  `json:"_id"`
		Type     int    // 1
		Name     string `json:"name"`
		Tag      string `json:"tag"`
		MarkDown string `json:"markdown"`
	}

	VideoSpace struct {
		Id     int64  `json:"_id"`
		Type   int    // 2
		Name   string `json:"name"`
		Tag    string `json:"tag"`
		Src    string `json:"src"`
		Format string `json:"format"`
	}

	ImageSpace struct {
		Id     int64  `json:"_id"`
		Type   int    // 3
		Name   string `json:"name"`
		Tag    string `json:"tag"`
		Src    string `json:"src"`
		Format string `json:"format"`
	}

	EmbeddingSpace struct {
		Id   int64  `json:"_id"`
		Type int    // 4
		Name string `json:"name"`
		Tag  string `json:"tag"`
		Src  string `json:"src"`
	}
)

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
type (
	UpdatedCommunityRequest Community
	UpdateSpaceRequest      Space

	SearchCommunityRequest struct {
		Name     string `json:"name"`
		Tag      string `json:"tag"`
		Pagesize int    `json:"pagesize"`
	}
	GetCommunityRequest struct {
		Name string `json:"name"`
		Tag  string `json:"tag"`
	}

	NewCommunityRequest struct {
		Name         string `json:"namespace_name"`
		Tag          string `json:"namespace_tag"`
		Display_name string `json:"name"`
		Description  string `json:"description"`
		Banner       string `json:"banner"`
	}
)

// response
type (
	GetCommunityResponse     Community
	SearchCommunityResponse  Community
	NewCommunityResponse     Community
	UpdatedCommunityResponse Community

	UpdateSpaceResponse Space
)
