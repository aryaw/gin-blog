package blog

import (
	"gin-blog/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.Use(middleware.JWTAuthMiddleware())
	v1.GET("/blog/hello", RenderBlogHello)
	v1.POST("/blog/create", CreateBlog)
	v1.POST("/blog/update/:id", UpdateBlog)
	v1.POST("/blog/delete/:id", DeleteBlog)

	v1.POST("/blog-category/create", CreateBlogCategory)
	v1.POST("/blog-category/update/:id", UpdateBlogCategory)
	v1.POST("/blog-category/delete/:id", DeleteBlogCategory)
}