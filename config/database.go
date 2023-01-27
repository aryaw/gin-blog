package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	var dbHOST = os.Getenv("DB_HOST")
	var dbPORT = os.Getenv("DB_PORT")
	var dbUSER = os.Getenv("DB_USER")
	var dbPASSWORD = os.Getenv("DB_PASSWORD")
	var dbNAME = os.Getenv("DB_NAME")

	dbConnection := dbUSER+":"+dbPASSWORD+"@tcp("+dbHOST+":"+dbPORT+")/"+dbNAME+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbConnection), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	Database, err := db.DB()
	Database.SetMaxIdleConns(10)
	Database.SetMaxOpenConns(100)
	Database.SetConnMaxLifetime(time.Hour)
	// DB = Database
	DB = db

	return DB
}

func GetDB() *gorm.DB {
	return DB
}