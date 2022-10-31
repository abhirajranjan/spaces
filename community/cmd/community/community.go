package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	app "github.com/abhirajranjan/spaces/community/internal/app"
	pb "github.com/abhirajranjan/spaces/proto/pkg/community"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommunityServiceServer(s, &app.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
