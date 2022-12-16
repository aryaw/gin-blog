package authcms

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (ctrl UserController) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "It's Wokrs"})
}