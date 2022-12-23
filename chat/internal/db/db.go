package db

import (
	"crypto/tls"
	"log"

	"github.com/abhirajranjan/spaces/chat/config"
	"github.com/abhirajranjan/spaces/chat/pkg/constants"
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

func executeMultiple(param *pb.QueryParameters, queries []string) (res *pb.Response, status *constants.Status) {
	if len(queries) == 0 {
		return nil, constants.Status_ErrZeroQuery
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
	res, _err := db.ExecuteBatch(batch)
	if _err != nil {
		logger.Logger.Sugar().Error(_err)
		return nil, constants.Status_ErrCql
	}
	return res, constants.Status_Ok
}

func execute(param *pb.QueryParameters, query string) (res *pb.Response, status *constants.Status) {
	exeQuery := &pb.Query{Cql: query, Parameters: param}
	res, _err := db.ExecuteQuery(exeQuery)
	if _err != nil {
		logger.Logger.Sugar().Error(_err)
		return nil, constants.Status_ErrCql
	}
	return res, constants.Status_Ok
}
