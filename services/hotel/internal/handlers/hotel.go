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
// @Router       /hotels/{id} [get]
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

// CreateHotel
// @Summary      Add new hotel
// @Description  Add new hotel
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Success		 201 {object} domain.Hotel
// @Param        request body domain.CreateHotelRequest true "Hotel data"
// @Router       /hotels/create [post]
func (h *HotelHandler) CreateHotel(c *gin.Context) {
	var hotel domain.CreateHotelRequest

	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.hotelService.Create(hotel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// UpdateHotel
// @Summary      Update hotel information
// @Description  Update hotel information
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Success		 200 {object} string
// @Param        id path int true "Hotel ID"
// @Param        request body domain.CreateHotelRequest true "Hotel data"
// @Router       /hotels/update/{id} [put]
func (h *HotelHandler) UpdateHotel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var hotel domain.CreateHotelRequest

	if err = c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.hotelService.Update(domain.ID(id), hotel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
