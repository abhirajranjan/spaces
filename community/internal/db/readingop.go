// Read operations
package db

import (
	pb "github.com/abhirajranjan/spaces/community/pkg/grpc"
)

// TODO: make getCommunities a stream compatible thing. Possibly a generator ?
// ! Currently doesnt stream data. cannot be integrated application api
func (d *db) GetCommunities(filter interface{}) ([]*pb.CommunityMetaData, error) {
	return filterQueries(filter)
}

func (d *db) GetCommunity(filter interface{}) (*pb.CommunityMetaData, error) {
	metadata, err := d.GetCommunities(filter)
	if err == nil {
		return new(pb.CommunityMetaData), err
	}
	// ? if only one metadata match then only return ?
	if len(metadata) == 1 {
		return metadata[0], err
	}
	if len(metadata) == 0 {
		return new(pb.CommunityMetaData), NoAccountExists
	}
	// TODO: what to do with multiple matched results ?
	return metadata[0], nil
}
