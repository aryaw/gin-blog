package main

import (
	"github.com/aryaw/gin-blog/config"
	"github.com/gin-gonic/gin"
)

func main() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// database
	DB := config.Init()
    h := config.New(DB)
}