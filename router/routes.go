package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadRoutes(r *gin.Engine)  {
	marsRoutes(r)

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello world!🤖 All your bases are belong to us!",
		})
	})
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})


}
