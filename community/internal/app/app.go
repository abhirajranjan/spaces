package app

import (
	"context"

	"github.com/abhirajranjan/spaces/community/internal/db"
	pb "github.com/abhirajranjan/spaces/community/pkg/grpc"
)

type Server struct {
	pb.UnimplementedCommunityServiceServer
}

func (s *Server) GetCommunity(ctx context.Context, in *pb.CommunityGetRequest) (*pb.CommunityMetaData, error) {
	return db.GetCommunity(in)
}

func (s *Server) CommunitySearch(in *pb.CommunityGetRequest, srv pb.CommunityService_CommunitySearchServer) error {
	return db.CommunitySearchStream(in, srv)
}
