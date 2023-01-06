package app

import (
	"fmt"
	"log"
	"os"
	"io"
	"path/filepath"
	"encoding/json"
	"net/http"

	"gin-blog/config"
	"gin-blog/middleware"
	"gin-blog/form"

	// module
	"gin-blog/app/authcms"
	"gin-blog/app/blog"
	
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type RouteItem struct {
	HttpMethod string
	AbsolutePath string
	// HandlerName string
}
var ListAvailRoutes []RouteItem

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

	// register custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mustalphanum", form.MustAlphaNum)
	}

	// database
	DB := config.Init()
    config.New(DB)
	
	modules := ReadModules("./app")
	fmt.Println(modules)

	InitModule(r, DB)

	// start go on env.port
	port := os.Getenv("PORT")
	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))
	
	
	// list all endpoint
	routesList := r.Routes()
	for _, rtLst := range routesList {
		rtItem := RouteItem {
			HttpMethod: rtLst.Method,
			AbsolutePath: rtLst.Path,
			// HandlerName: rtLst.Handler,
		}
		ListAvailRoutes = append(ListAvailRoutes, rtItem)
	}

	endPointList := r.Group("/v1")
    endPointList.GET("/endpoint", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"endpoint": ListAvailRoutes,
		})
	})

	// register middleware
	r.Use(middleware.BodySizeMiddleware())
	r.Use(middleware.JWTAuthMiddleware())

	fmt.Println("====================================================")
	fmt.Println("")
	fmt.Println("====================================================")
	r.Run(":" + port)
}



func InitModule(r *gin.Engine, DB *gorm.DB) {
	// module routes
	authcms.Routes(r)
	blog.Routes(r)
	// discovery.Routes(r)

	// module migration
	authcms.Migrate(DB)
	blog.Migrate(DB)
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