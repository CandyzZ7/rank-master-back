package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"github.com/itchyny/base58-go"
)

func GetSnowflakeID() string {
	node, _ := snowflake.NewNode(1)
	// Generate a snowflake ID.
	id := node.Generate()
	// Convert snowflake ID to base58
	enc := base58.BitcoinEncoding
	shortID, _ := enc.Encode(id.Bytes())

	return string(shortID)
}
