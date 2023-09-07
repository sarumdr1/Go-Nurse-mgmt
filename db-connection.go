package main

import (
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

var db *gorm.DB
var err error

func DataMigration() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  password := os.Getenv("PASSWORD")
  userName := os.Getenv("USER_NAME")
  dbName := os.Getenv("DBNAME")

  connStr := "host=localhost user="+userName+" password="+password+" dbname="+dbName+" sslmode=disable"
   
    db, err = gorm.Open(postgres.Open(connStr),&gorm.Config{})
    if err != nil {
      panic(err)
    } else{
        log.Println("DB connected")
    }
    db.AutoMigrate(&Nurse{})
}