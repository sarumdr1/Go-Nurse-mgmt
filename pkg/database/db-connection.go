package database

import (
	"example/go-nurse-mgmt/pkg/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func DataMigration() {

	password := os.Getenv("PASSWORD")
	userName := os.Getenv("USER_NAME")
	dbName := os.Getenv("DB_NAME")

	connStr := "host=localhost user=" + userName + " password=" + password + " dbname=" + dbName + " sslmode=disable"

	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Println("DB connected")
	}
	db.AutoMigrate(&models.Nurse{})
	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return db
}
