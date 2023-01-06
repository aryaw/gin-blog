package blog

import (
	// "gin-blog/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	v1 := r.Group("/v1")
	// v1.Use(middleware.JWTAuthMiddleware())
	v1.GET("/blog/hello", RenderBlogHello)
	v1.POST("/blog/create", CreateBlog)
	v1.POST("/blog/update/:id", UpdateBlog)
}