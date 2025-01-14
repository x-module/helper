/**
* Created by GoLand
* @file mysql.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/1/27 11:42 上午
* @desc 初始化管理后台数据库
 */

package dirver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	DbType    = "mysql"
)

var DB *gorm.DB

type LinkParams struct {
	Host        string
	Port        int
	UserName    string
	DbName      string
	Password    string
	MaxOpenConn int
	MaxIdleConn int
	LogLevel    int
}

// InitializeDB 初始化管理后台数据库
func InitializeDB(params LinkParams) (*gorm.DB, error) {
	linkParams := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(linkParams, params.UserName, params.Password, params.Host, params.Port, params.DbName)
	newLogger := logger.New(
		log.New(NewLogger(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,                      // Slow SQL threshold
			LogLevel:                  logger.LogLevel(params.LogLevel), // Log level
			IgnoreRecordNotFoundError: true,                             // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,                            // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	// 链接池设置
	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(params.MaxOpenConn)
	sqlDb.SetMaxIdleConns(params.MaxIdleConn)
	// db.LogMode(params.Mode == DebugMode)
	// db.LogMode(false)
	return db, nil
}

type Logger struct {
}

func NewLogger() *Logger {
	return new(Logger)
}

// Write 实现Write接口，用于写入
func (l *Logger) Write(p []byte) (n int, err error) {
	return 1, nil
}
