package user

import (
	"goto/database/mysql"
)

func Migration() {
	// 迁移表结构
	migrates := make(map[string]interface{})
	migrates["user"] = User{}
	mysql.AutoMigrate(migrates)
}
