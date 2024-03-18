package server

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	mysqlConn *gorm.DB
	dsn       string
	debug     logger.LogLevel
)

func init() {
	dsn = os.Getenv("GOLANG_SERVER_DSN_STR")
	if dsn == "" {
		log.Fatalf("环境变量设置有误，无法启动服务DSN")
	}
	switch os.Getenv("GOLANG_SERVER_MYSQL_DEBUG") {
	case "info":
		debug = logger.Info
	case "silent":
		debug = logger.Silent
	case "warn":
		debug = logger.Warn
	default:
		debug = logger.Error
	}

	dbPoolConfig := mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	})
	gormConfig := gorm.Config{}
	db, err := gorm.Open(dbPoolConfig, &gormConfig, &gorm.Config{Logger: logger.Default.LogMode(debug)})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	mysqlConn = db
}

func MSlConn() *gorm.DB {
	return mysqlConn
}

type MysqlTable interface {
	TableName() string
	Conn() *gorm.DB
	First() error
}

// Mysql表格注册
func RegisterMysqlTable(t ...interface{}) {
	// 自动创建表格，如果表格已经存在，检查字段是否有变化
	if err := mysqlConn.AutoMigrate(t...); err != nil {
		log.Fatalf("表格初始化失败：%v", err)
	}
}

func ServerDatabaseClose() {
	db, _ := mysqlConn.DB()
	_ = db.Close()
}
