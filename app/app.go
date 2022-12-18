package app

import (
	"fmt"
	"log"
	"os"
	"io"
	"path/filepath"
	"encoding/json"

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

	// database
	DB := config.Init()
    config.New(DB)
	

	ReadPackage("")
	
	// module routes
	authcms.Routes(r)
	// module migration
	authcms.Migrate(DB)

	// start go on ev.port
	port := os.Getenv("PORT")
	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))
	fmt.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))
	r.Run(":" + port)
	r.Routes()
}

func ReadPackage(dir string) {
	// read the package
	appDir := "./app"
	var modulePath string

	if dir == "" {
		dir = appDir
	}

	f, err := os.Open(dir)
    if err != nil {
        fmt.Println(err)
        return
    }
    files, err := f.Readdir(0)
    if err != nil {
        fmt.Println(err)
        return
    }

    for _, v := range files {
		if (v.IsDir()) {
			dir = appDir+"/"+v.Name()
			ReadPackage(dir)
		}

		extension := filepath.Ext(v.Name())
		if extension == ".json" {
			modulePath = dir+"/"+v.Name()
			GetPackageInfo(modulePath)
		}
    }
}

type Module struct {
    Name   string `json:"name"`
    Title   string `json:"title"`
    Version   string `json:"version"`
    Description   string `json:"description"`
    Author   string `json:"author"`
    Email   string `json:"email"`
}

func GetPackageInfo(path string) {
	jsonFile, err := os.Open(path)
    if err != nil {
        fmt.Println(err)
    }

	byteValue, _ := io.ReadAll(jsonFile)
	var module Module
	json.Unmarshal(byteValue, &module)
	fmt.Println(module.Name)
}