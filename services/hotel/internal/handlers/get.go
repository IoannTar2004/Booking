package handlers

import "github.com/gin-gonic/gin"

func getHotelById(c *gin.Context) {
	c.JSON(200, gin.H{"status": "hotel"})
}
