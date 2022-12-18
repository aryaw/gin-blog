package app

import (
	"fmt"
	"log"
	"os"

	"gin-blog/config"
	// "gin-blog/middleware"
	"gin-blog/app/authcms"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	// r.Use(middleware.CORSMiddleware())
	
	// register routes
	authcms.Routes(r)

	// database
	DB := config.Init()
    config.New(DB)
	authcms.Migrate(DB)

	// start go on ev.port
	port := os.Getenv("PORT")
	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))
	fmt.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))
	r.Run(":" + port)
	r.Routes()
}