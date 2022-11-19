// write operations
package db

import (
	pb "github.com/abhirajranjan/spaces/community/pkg/grpc"
)

func NewCommunity(metadata *pb.CommunityMetaData) (Hex, error) {
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
	_, err := filterQueriesFindOne(&filterNameTags)
	switch err {
	// true if no account matches data
	case NoAccountExists:
		return nil
	//if error is nil then account exists
	case nil:
		return AccountAlreadyExists
	// if any other error occured returns DBError
	default:
		handleError("DBError(verifyNewMetaData)", err)
		return DBError
	}
}
