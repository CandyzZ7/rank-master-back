package snowflake

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

var mutex sync.Mutex

// 写死node【注：优化】
var node int64 = 1

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
