package snowflake

import (
	"sync"

	"github.com/bwmarrin/snowflake"

	"rank-master-back/internal/config"
)

var (
	mutex sync.Mutex
	node  int64
)

func InitNode(c config.Config) {
	if c.WorkerId == 0 {
		node = 1
	}
	node = c.WorkerId
}

func newNode() (*snowflake.Node, error) {
	node, err := snowflake.NewNode(node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func GenerateSnowflakeID() snowflake.ID {
	mutex.Lock()
	defer mutex.Unlock()
	node, _ := newNode()
	// Generate a snowflake ID.
	id := node.Generate()
	return id
}

func GenerateDefaultSnowflakeID() string {
	mutex.Lock()
	defer mutex.Unlock()
	// 忽略错误
	node, _ := newNode()
	// Generate a snowflake ID.
	id := node.Generate().Base58()
	return id
}
