package db

import (
	"context"
	"log"
	"reflect"
	"strings"

	pb "github.com/abhirajranjan/spaces/community/pkg/grpc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.Background()

// initialise the database
func InitDb(database string, col string) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(database).Collection(col)
}

// TODO: getrequest to bson_M
// ? need to convert all key to lowercase? can't just lowercase the first ?
func metadataToBson_M(buffer interface{}) bson.M {
	switch b := buffer.(type) {
	case bson.M:
		return b
	}
	var bsonM bson.M = make(map[string]interface{})

	s := reflect.ValueOf(buffer).Elem()
	typeOfBuffer := s.Type()

	// start from 3 coz 0, 1, 2 are occupied by grpc defaults
	for i := 3; i < s.NumField(); i++ {
		//* fmt.Printf("%d: %s %s = %v\n", i, typeOfBuffer.Field(i).Name, f.Type(), s.FieldByName(typeOfBuffer.Field(i).Name))

		switch f := s.Field(i); f.Kind() {
		case reflect.String:
			val := s.FieldByName(typeOfBuffer.Field(i).Name)
			if !val.CanConvert(reflect.TypeOf("")) || val.String() == "" {
				break
			}
			bsonM[strings.ToLower(typeOfBuffer.Field(i).Name)] = s.FieldByName(typeOfBuffer.Field(i).Name).String()

		case reflect.Int:
			if !s.FieldByName(typeOfBuffer.Field(i).Name).CanUint() {
				break
			}
			bsonM[strings.ToLower(typeOfBuffer.Field(i).Name)] = s.FieldByName(typeOfBuffer.Field(i).Name).String()

		// check for all slices and initialize its type
		case reflect.Slice:
			// only initialize if it has value
			if f.Len() == 0 {
				break
			}

			arr := []map[string]interface{}{}
			//fmt.Println(reflect.TypeOf(bsonM[name]))
			for i := 0; i < f.Len(); i++ {
				// https://stackoverflow.com/questions/25102827/struct-value-of-pointer-array-and-slice
				g := f.Index(i).Elem()
				typeOfG := g.Type()

				for j := 3; j < g.NumField(); j++ {
					mapping := map[string]interface{}{
						strings.ToLower(typeOfG.Field(j).Name): g.FieldByName(typeOfG.Field(j).Name).String(),
					}
					arr = append(arr, mapping)
				}
				bsonM[strings.ToLower(typeOfBuffer.Field(i).Name)] = arr
			}
		}
	}

	return bsonM
}

// convert filter to bson.M type
func preProcessRequest(filter interface{}) bson.M {
	return metadataToBson_M(filter)
}

// stream the result to stream.Send() and returns error
// TODO: set options for better matching of incomplete query
func streamFilterQueries(filter interface{}, stream interface {
	Send(*pb.CommunityMetaData) error
}) chan error {
	c := make(chan error, 1)
	go func() {
		cur, err := collection.Find(ctx, preProcessRequest(filter))
		if err != nil {
			c <- err
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			var m pb.CommunityMetaData
			if err := cur.Decode(&m); err != nil {
				c <- err
			}
			if err := stream.Send(&m); err != nil {
				c <- err
			}
		}
		if err := cur.Err(); err != nil {
			c <- err
		}
		close(c)
	}()
	return c
}

/* find one matched result from db
returns nil and error equals NoAccountsExists if no match found
returns matched document and error equals nil if match found
*/
// TODO: set options for better matching of incomplete query
func filterQueriesFindOne(filter interface{}) (*pb.CommunityMetaData, error) {
	doc := collection.FindOne(ctx, preProcessRequest(filter))
	if doc.Err() == mongo.ErrNoDocuments {
		return nil, NoAccountExists
	}
	var m pb.CommunityMetaData
	if err := doc.Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

// insert Data to the document
func insertData(metadata *pb.CommunityMetaData) (Hex, error) {
	res, err := collection.InsertOne(ctx, metadata)
	if err != nil {
		return Hex{value: "0"}, err
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return Hex{value: oid.Hex()}, nil
	}
	return Hex{value: "0"}, err
}
