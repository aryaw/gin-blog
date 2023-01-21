package authcms

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/", Hello)
	v1.GET("/login", RenderLogin)
	v1.POST("/login", PostLogin)
	v1.POST("/register", PostRegister)
}

// func (app *Application) Routes() {
// 	app.Router := gin.Default()
// 	v1 := app.Router.Group("/v1")
// 	v1.GET("/", Hello)
// 	v1.GET("/login", RenderLogin)
// 	v1.POST("/login", PostLogin)
// 	v1.POST("/register", PostRegister)
// }