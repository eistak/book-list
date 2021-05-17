package driver

import (
	"books-list/utils"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB connects postgresql by GORM
func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.LogFatal(err)

	return db
}
