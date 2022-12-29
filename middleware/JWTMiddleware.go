package middleware

import (
    "net/http"

    "gin-blog/app/authcms"
    "github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(context *gin.Context) {
        err := authcms.ValidateJWT(context)
        if err != nil {
            context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
            context.Abort()
            return
        }
        context.Next()
    }
}