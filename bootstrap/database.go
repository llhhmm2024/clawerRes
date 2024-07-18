package bootstrap

import "crawler/pkg/database"

// SetupDB 初始化数据库
func SetupDB() {
	database.InitGorm()
	go database.MysqlPipeline()
}
