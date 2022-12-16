package authcms

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthCmsController struct{}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "It's Wokrs"})
}