package authcms

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthCmsController struct{}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "It's Wokrs"})
}

func RenderLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "render Login"})
}

func PostRegister(context *gin.Context) {
    var input AuthenticationInput

    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error request": err.Error()})
        return
    }
    
    user := UserModel{
        FirstName: input.FirstName,
        LastName: input.LastName,
        Email: input.Email,
        Password: input.Password,
    }
    
    savedUser, err := user.Save()

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}