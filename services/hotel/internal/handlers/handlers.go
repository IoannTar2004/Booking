package handlers

import (
	"hotel/internal/infrastructure/database"
	"hotel/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Init(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	hotelRepo := database.NewHotelRepository(db)
	hotelService := services.NewHotelService(hotelRepo)
	hotelHandler := NewHotelHandler(hotelService)

	hotel := router.Group("/api/v1/hotels")
	hotel.GET("/:id", hotelHandler.GetHotelById)
	hotel.GET("/all", hotelHandler.GetHotels)
	hotel.POST("/create", hotelHandler.CreateHotel)
	hotel.PUT("/update/:id", hotelHandler.UpdateHotel)

	return router
}
