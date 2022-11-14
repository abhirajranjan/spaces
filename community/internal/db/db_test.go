package db

import (
	"testing"

	pb "github.com/abhirajranjan/spaces/community/pkg/grpc"
	"github.com/stretchr/testify/assert"
)

var d = db{}

func TestMain(m *testing.M) {
	d.Init("collection", "metadata")
	m.Run()
}

func TestFilterConvertor(t *testing.T) {
	assert := assert.New(t)
	buffer := pb.CommunityMetaData{
		Name: "spaces talks",
		Tag:  "gamers",
	}

	data := map[string]interface{}{
		"name": "spaces talks",
		"tag":  "gamers",
	}

	bsonM, errM := filterQueries(&buffer)
	bson, err := _filterQueries(&data)
	assert.ElementsMatch(bson, bsonM)
	assert.ErrorIs(err, errM)

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
		_, err := d.NewCommunity(&tests[i].input)
		assert.ErrorIs(err, tests[i].expected)
	}
}
