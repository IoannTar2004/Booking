package handlers

import (
	"errors"
	"hotel/internal/domain"
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
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success		 200 {object} domain.Hotel
// @Param        id   path      int  true  "Hotel ID"
// @Router       /hotels/hotel/{id} [get]
func (h *HotelHandler) GetHotelById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	hotel, err := h.hotelService.GetHotelById(id)

	if errors.Is(err, domain.HotelNotFoundError{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, hotel)
	}
}

// GetHotels
// @Summary      Get list of hotels
// @Description  Get list of hotels
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success		 200 {object} []domain.Hotel
// @Param        name query string false "Hotel name"
// @Param        city query string false "City location"
// @Param        stars query int false "Stars"
// @Param        limit query int true "Limit" default(5)
// @Param        offset query int true "Offset" default(0)
// @Router       /hotels/all [get]
func (h *HotelHandler) GetHotels(c *gin.Context) {
	var getHotelsRequest domain.GetHotelsRequest

	if err := c.ShouldBindQuery(&getHotelsRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotels, err := h.hotelService.GetAll(getHotelsRequest)
	if errors.Is(err, domain.HotelNotFoundError{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, hotels)
	}
}
