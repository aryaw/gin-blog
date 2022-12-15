package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) DbHandler {
	return DbHandler{db}
}

func Init() *gorm.DB {
	var dbHOST = os.Getenv("DB_HOST")
	var dbPORT = os.Getenv("DB_PORT")
	var dbUSER = os.Getenv("DB_USER")
	var dbPASSWORD = os.Getenv("DB_PASSWORD")
	var dbNAME = os.Getenv("DB_NAME")


	dbConnection := dbUSER":"+dbPASSWORD+"@tcp("+dbHOST+":"+dbPORT+")/"+dbNAME+"?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dbConnection), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
