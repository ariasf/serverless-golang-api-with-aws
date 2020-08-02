package mars

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMarsSolsWeatherInfoController(ctx *gin.Context) {

	response, err := GetAllSolsWeather()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
