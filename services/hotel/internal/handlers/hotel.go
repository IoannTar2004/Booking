package handlers

import (
	"hotel/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HotelHandler struct {
	hotelService *services.HotelService
}

func NewHotelHandler(hotelService *services.HotelService) *HotelHandler {
	return &HotelHandler{hotelService: hotelService}
}

// GetHotelById
// @Summary      Get hotel by ID
// @Description  Get hotel information by ID
// @Tags         hotels
// @Accept       json
// @Produce      json
// @Success		 200 {object} domain.Hotel
// @Param        id   path      int  true  "Hotel ID"
// @Router       /hotels/hotel/{id} [get]
func (h *HotelHandler) GetHotelById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	hotel, err := h.hotelService.GetHotelById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, hotel)
}
