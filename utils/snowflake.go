package utils

import (
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
	once sync.Once
)

// initNode 初始化雪花节点
func initNode() {
	var err error
	once.Do(func() {
		// 创建一个新的雪花节点，使用当前时间的秒数作为节点ID
		// 这样可以在一定程度上保证节点ID的唯一性
		nodeID := time.Now().UnixNano() % 1024
		node, err = snowflake.NewNode(nodeID)
		if err != nil {
			panic(err)
		}
	})
}

// GenerateSnowflakeID 生成雪花ID
func GenerateSnowflakeID() string {
	initNode()
	return node.Generate().String()
}

// GenerateSnowflakeIDInt64 生成雪花ID的int64形式
func GenerateSnowflakeIDInt64() int64 {
	initNode()
	return node.Generate().Int64()
}
