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

type db struct{}

func (d *db) Init(database string, col string) {
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

// TODO: break function for robust and schema less conversion
// ? need to convert all key to lowercase? can't just lowercase the first ?
func metadataToBson_M(buffer interface{}) bson.M {
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

// TODO: make fiterQueries a stream compatible thing. Possibly a generator ?
// filter queries from database
func filterQueries(filter interface{}) ([]*pb.CommunityMetaData, error) {
	return _filterQueries(metadataToBson_M(filter))
}

func _filterQueries(filter interface{}) ([]*pb.CommunityMetaData, error) {
	var metadata []*pb.CommunityMetaData

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return metadata, err
	}
	defer cur.Close(ctx)

	// iterate over all queries in cur and add it to metadata
	for cur.Next(ctx) {
		var m pb.CommunityMetaData
		if err := cur.Decode(&m); err != nil {
			return metadata, err
		}
		metadata = append(metadata, &m)
	}

	if err := cur.Err(); err != nil {
		return metadata, err
	}

	// check if no document find with given filter
	if len(metadata) == 0 {
		return metadata, NoAccountExists
	}
	return metadata, nil
}

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
