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
    var input RegisterInput

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

func PostLogin(context *gin.Context) {
	var input AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usermodel, err := FindUserByEmail(input.Email)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = usermodel.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := GenerateJWT(usermodel)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
} 