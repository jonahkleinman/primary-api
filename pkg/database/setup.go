package database

import (
	"fmt"
	"github.com/VATUSA/primary-api/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func Connect(dbConfig *config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)

	logLevel := logger.Info
	switch dbConfig.LoggerLevel {
	case "silent":
		logLevel = logger.Silent
		break
	case "error":
		logLevel = logger.Error
		break
	case "warn":
		logLevel = logger.Warn
		break
	}

	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatal("[Database] Connection Error:", err)
	}

	return db
}
