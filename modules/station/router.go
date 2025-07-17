package station

import (
	"net/http"

	"github.com/Dito-7/mrt-schedules-golang/common/response"
	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/stations")
	station.GET("", func(ctx *gin.Context) {
		GetAllStations(ctx, stationService)
	})
}

func GetAllStations(ctx *gin.Context, service Service) {
	datas, err := service.GetAllStations()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Successfully Get All stations",
		Data:    datas,
	})
}
