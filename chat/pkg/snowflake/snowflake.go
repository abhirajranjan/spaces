package snowflake

import (
	"sync"

	"github.com/abhirajranjan/spaces/chat/pkg/logger"
	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node
var mutex = &sync.RWMutex{}

type ID = snowflake.ID

func init() {
	// Create a new Node with a Node number of 1
	var err error
	mutex.Lock()
	defer mutex.Unlock()
	node, err = snowflake.NewNode(1)
	if err != nil {
		logger.Logger.Sugar().Error(err)
	}
}

func Generate() snowflake.ID {
	mutex.RLock()
	defer mutex.RUnlock()
	return node.Generate()
}

func ParseInt64(val int64) *snowflake.ID {
	a := snowflake.ParseInt64(val)
	return &a
}
