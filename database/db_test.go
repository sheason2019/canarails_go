package database_test

import (
	"testing"

	"canarails.dev/database"
)

// 测试数据库连接和数据库迁移
func TestConnectDb(t *testing.T) {
	database.GetDb()
}
