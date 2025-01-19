package db

import (
	"os"
	logger "server/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	var err error
	dsn := os.Getenv("DB_STRING")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Critical(err, "Failed to connect to database")
	} else {
		logger.Info("Connected to database")
	}
}
