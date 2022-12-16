package authcms

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/", Hello)
}