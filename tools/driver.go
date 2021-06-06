package tools

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB connects postgresql by GORM
func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	LogFatal(err)

	return db
}
