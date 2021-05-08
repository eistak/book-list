package driver

import (
	"books-list/utils"
	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

// ConnectDB connects postgresql by GORM
func ConnectDB() *gorm.DB {
	err := gotenv.Load()
	utils.LogFatal(err)

	dsn := os.Getenv("DSN")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.LogFatal(err)

	return db
}
