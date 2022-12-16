package authcms

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	v1 := route.Group("/v1") {
		v1.GET("/", authcms.Hello)
	}
}