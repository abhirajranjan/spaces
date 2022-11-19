package db

import (
	"log"
	"testing"

	pb "github.com/abhirajranjan/spaces/community/pkg/grpc"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMain(m *testing.M) {
	InitDb("collection", "metadata")
	m.Run()
}

type streamer struct {
	Slice1 []*pb.CommunityMetaData
}

func (s streamer) Send(i *pb.CommunityMetaData) error {
	s.Slice1 = append(s.Slice1, i)
	return nil
}

func TestFilterQueriesFindOne(t *testing.T) {
	buffer := &pb.CommunityMetaData{
		Name: "spaces talks",
		Tag:  "gamers",
	}

	ds, err := filterQueriesFindOne(buffer)
	if err != nil {
		t.Log(err)
	}
	t.Logf("findone: %v", ds)
}

func TestStreamFilterQueries(t *testing.T) {
	buffer := &pb.CommunityMetaData{
		Name: "spaces talks",
		Tag:  "gamers",
	}
	var data bson.M = map[string]interface{}{
		"name": "spaces talks",
		"tag":  "gamers",
	}

	s1 := streamer{}
	s2 := streamer{}
	channelerr1 := streamFilterQueries(buffer, s1)
	channelerr2 := streamFilterQueries(data, s2)

	if err, open := <-channelerr1; open {
		t.Fail()
		log.Println(err)
	}

	if err, open := <-channelerr2; open {
		t.Fail()
		log.Println(err)
	}

	assert.Equal(t, s1.Slice1, s2.Slice1)
}

func TestNewCommunity(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input    pb.CommunityMetaData
		expected error
	}{
		{
			input: pb.CommunityMetaData{
				Id: 1,
			},
			expected: NameCannotBeNull,
		},
		{
			input: pb.CommunityMetaData{
				Name:        "",
				Tag:         "",
				Description: "",
			},
			expected: NameCannotBeNull,
		},
		{
			input: pb.CommunityMetaData{
				Name: "hello",
			},
			expected: TagCannotBeNull,
		},
		{
			input: pb.CommunityMetaData{
				Name: "hello",
				Tag:  "1212",
			},
			expected: DescCannotBeNull,
		},
	}
	for i := 0; i < len(tests); i++ {
		_, err := NewCommunity(&tests[i].input)
		assert.ErrorIs(err, tests[i].expected)
	}
}
