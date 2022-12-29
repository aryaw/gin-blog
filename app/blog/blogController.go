package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderBlogHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "render Blog Hello"})
}
