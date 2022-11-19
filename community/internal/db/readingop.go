// Read operations
package db

import (
	pb "github.com/abhirajranjan/spaces/community/pkg/grpc"
)

func GetCommunity(filter *pb.CommunityGetRequest) (*pb.CommunityMetaData, error) {
	result, err := filterQueriesFindOne(filter)
	switch err {
	case NoAccountExists:
		return nil, err
	case nil:
		return result, nil
	default:
		//* on unmarshal failed, return no account exists and log err
		handleError("dataUnmarshalError (GetCommunity)", err)
		return nil, NoAccountExists
	}
}

func CommunitySearchStream(in *pb.CommunityGetRequest, stream interface {
	Send(*pb.CommunityMetaData) error
}) error {
	errorChannel := streamFilterQueries(in, stream)
	if err, open := <-errorChannel; open {
		handleError("dbSearchStream", err)
		close(errorChannel)
		return err
	}
	return nil
}
