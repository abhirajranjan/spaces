package db

import (
	"crypto/tls"
	"log"

	"github.com/abhirajranjan/spaces/chat/config"
	"github.com/abhirajranjan/spaces/chat/pkg/logger"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/auth"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/client"
	pb "github.com/stargate/stargate-grpc-go-client/stargate/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var db *client.StargateClient

func init() {
	// Create connection with authentication
	// For Astra DB:
	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
	}

	conn, err := grpc.Dial(config.DbURI(), grpc.WithTransportCredentials(credentials.NewTLS(tlsconfig)),
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(
			auth.NewStaticTokenProvider(config.DbToken()),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	db, err = client.NewStargateClientWithConn(conn)

	if err != nil {
		log.Fatalf("error creating client %v", err)
	}
}

func executeMultiple(queries ...string) *pb.Response {
	if len(queries) == 0 {
		return nil
	}
	var batchqueries []*pb.BatchQuery

	for i := range queries {
		tempquery := pb.BatchQuery{Cql: queries[i]}
		batchqueries = append(batchqueries, &tempquery)
	}

	batch := &pb.Batch{
		Type:    pb.Batch_LOGGED,
		Queries: batchqueries,
	}
	res, err := db.ExecuteBatch(batch)
	if err != nil {
		logger.Logger.Sugar().Error(err)
	}
	return res
}

func execute(param *pb.QueryParameters, query string) (*pb.Response, error) {
	exeQuery := &pb.Query{Cql: query, Parameters: param}
	res, err := db.ExecuteQuery(exeQuery)
	if err != nil {
		logger.Logger.Sugar().Error(err)
	}
	return res, ErrCql
}
