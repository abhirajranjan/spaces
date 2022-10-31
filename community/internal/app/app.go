package app

import (
	"context"
	"log"
	"sync"

	pb "github.com/abhirajranjan/spaces/proto/pkg/community"
)

type Server struct {
	pb.UnimplementedCommunityServiceServer
}

func (s *Server) GetCommunity(ctx context.Context, in *pb.CommunityGetRequest) (*pb.CommunityMetaData, error) {
	log.Printf("Received: %v", in.RequestType)
	return &pb.CommunityMetaData{Status: pb.STATUS_OK.Enum()}, nil
}

func (s *Server) CommunitySearch(in *pb.CommunityGetRequest, srv pb.CommunityService_CommunitySearchServer) error {
	var wg = new(sync.WaitGroup)
	mutex := new(sync.Mutex)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(count int64) {
			mutex.Lock()
			defer wg.Done()
			desc := "Student"
			resp := pb.CommunityMetaData{
				Id:          "AR",
				Name:        "Abhiraj Ranjan",
				Description: &desc,
			}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			mutex.Unlock()
			log.Printf("finishing request number : %d", count)
		}(int64(i))
	}
	wg.Wait()
	return nil
}
