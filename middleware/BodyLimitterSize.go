package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BodySizeMiddleware()  gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var maxBytes int64 = 1024 * 1024 * 2 // 2MB

		var wtr http.ResponseWriter = ctx.Writer
		ctx.Request.Body = http.MaxBytesReader(wtr, ctx.Request.Body, maxBytes)

		ctx.Next()
	}
}