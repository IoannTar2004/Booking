package handlers

import "github.com/gin-gonic/gin"

func InitHandlers() *gin.Engine {
	router := gin.Default()
	hotel := router.Group("/hotels")
	hotel.GET("/hotel/:id", getHotelById)

	return router
}
