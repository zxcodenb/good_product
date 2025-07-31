package utils

import (
	"testing"
)

func TestGenerateSnowflakeID(t *testing.T) {
	// 测试生成雪花ID
	id := GenerateSnowflakeID()
	if id == "" {
		t.Error("Expected a non-empty ID, but got an empty string")
	}

	// 测试生成的ID是否为数字字符串
	for _, char := range id {
		if char < '0' || char > '9' {
			t.Errorf("Expected a numeric string, but got: %s", id)
		}
	}
}

func TestGenerateSnowflakeIDInt64(t *testing.T) {
	// 测试生成雪花ID的int64形式
	id := GenerateSnowflakeIDInt64()
	if id <= 0 {
		t.Errorf("Expected a positive int64, but got: %d", id)
	}
}

func TestUniqueIDs(t *testing.T) {
	// 测试生成的ID是否唯一
	id1 := GenerateSnowflakeID()
	id2 := GenerateSnowflakeID()
	
	if id1 == id2 {
		t.Errorf("Expected unique IDs, but got duplicates: %s", id1)
	}
}