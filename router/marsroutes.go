package router

import (
	"github.com/gin-gonic/gin"
	"serverless-golang-api-with-aws/mars"
)

func marsRoutes(r *gin.Engine) {
	marsRoutes := r.Group("/mars")
	{
		marsRoutes.GET("/weather/sols", func(context *gin.Context) {
			mars.GetMarsSolsWeatherInfoController(context)
		})
	}

}
