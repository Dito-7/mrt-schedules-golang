package station

import (
	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/station")
	station.GET("", func(ctx *gin.Context) {
		GetAllStations(ctx, stationService)
	})
}

func GetAllStations(ctx *gin.Context, service Service) {

}
