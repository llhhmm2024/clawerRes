package database

import (
	"crawler/configs"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func InitGorm() {
	db = gormMysql()
}

func gormMysql() *gorm.DB {
	m := configs.Cfg.Database
	// dsn := m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		//DriverName:                "mysql",
		DSN:                       m.Dsn(),
		Conn:                      nil,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.Debug, m.DryRun)); err != nil {
		log.Fatalf("Mysql 启动异常")
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetConnMaxIdleTime(8 * time.Minute) // 设置最大空闲时间。默认值为0，表示不超时。
		sqlDB.SetConnMaxLifetime(60 * time.Second)
		sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，默认值为2。如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
		sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数,默认值为0，表示没有限制。

		return db
	}
}

func gormConfig(mod, dry_run bool) *gorm.Config {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		})
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		DryRun:         dry_run,
	}
	if mod {
		config.Logger = newLogger
	}
	return config
}
