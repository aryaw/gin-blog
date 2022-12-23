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
	"gorm.io/gorm"
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
	
	modules := ReadModules("./app")
	fmt.Println(modules)

	InitModule(r, DB)

	// start go on ev.port
	port := os.Getenv("PORT")
	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))
	fmt.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))
	r.Run(":" + port)
	r.Routes()
}

func InitModule(r *gin.Engine, DB *gorm.DB) {
	// module routes
	authcms.Routes(r)

	// module migration
	authcms.Migrate(DB)
}

var Modules []string
func ReadModules(dir string) []string {
	// read the module
	appDir := "./app"
	var modulePath string

	f, err := os.Open(dir)
    if err != nil {
        fmt.Println(err)
    }
    files, err := f.Readdir(0)
    if err != nil {
        fmt.Println(err)
    }

	for _, v := range files {
		if (v.IsDir()) {
			Modules = append(Modules, v.Name())
			dir = appDir+"/"+v.Name()
			ReadModules(dir)
		}

		extension := filepath.Ext(v.Name())
		if extension == ".json" {
			modulePath = dir+"/"+v.Name()
			GetPackageInfo(modulePath)
		}
    }

	return Modules
}

type ModuleInfo struct {
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
	var moduleinfo ModuleInfo
	json.Unmarshal(byteValue, &moduleinfo)
	// fmt.Println(moduleinfo.Name)
}