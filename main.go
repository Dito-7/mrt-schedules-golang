package main

import (
	"github.com/Dito-7/mrt-schedules-golang/modules/station"
	"github.com/gin-gonic/gin"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {
	var (
		router = gin.Default()
		api    = router.Group("/v1/api")
	)

	station.Initiate(api)

	router.Run(":8080")

}
