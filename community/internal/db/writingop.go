// write operations
package db

import (
	pb "github.com/abhirajranjan/spaces/community/pkg/grpc"
)

func (d *db) NewCommunity(metadata *pb.CommunityMetaData) (Hex, error) {
	err := verifyNewMetaData(metadata)
	if err != nil {
		return Hex{value: "0"}, err
	}

	hex, err := insertData(metadata)
	return hex, err
}

// verification of new Community to be added
func verifyNewMetaData(metadata *pb.CommunityMetaData) error {
	if metadata.Name == "" {
		return NameCannotBeNull
	}
	if metadata.Tag == "" {
		return TagCannotBeNull
	}
	if metadata.Description == "" {
		return DescCannotBeNull
	}
	filterNameTags := pb.CommunityMetaData{
		Name: metadata.Name,
		Tag:  metadata.Tag,
	}
	ds, err := filterQueries(&filterNameTags)
	if err == NoAccountExists && len(ds) == 0 {
		return nil
	}
	return AccountAlreadyExists
}
