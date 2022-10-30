package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/abhirajranjan/spaces/microservices/community/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedCommunityServiceServer
}

func (s *server) GetCommunity(ctx context.Context, in *pb.CommunityGetRequest) (*pb.CommunityMetaData, error) {
	log.Printf("Received: %v", in.RequestType)
	return &pb.CommunityMetaData{Status: pb.STATUS_OK.Enum()}, nil
}

func (s *server) CommunitySearch(in *pb.CommunityGetRequest, srv pb.CommunityService_CommunitySearchServer) error {
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

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommunityServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
