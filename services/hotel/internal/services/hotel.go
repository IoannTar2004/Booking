package services

import (
	"hotel/internal/domain"
	"hotel/internal/repositories"
)

type HotelService struct {
	hotelRepo repositories.Hotel
}

func NewHotelService(hotelRepo repositories.Hotel) *HotelService {
	return &HotelService{hotelRepo: hotelRepo}
}

func (h *HotelService) GetHotelById(id int64) (*domain.Hotel, error) {
	return h.hotelRepo.GetById(id)
}

func (h *HotelService) GetAll(filter domain.GetHotelsRequest) ([]domain.Hotel, error) {
	return h.hotelRepo.GetAll(filter)
}

func (h *HotelService) Create(createHotelRequest domain.CreateHotelRequest) (domain.ID, error) {
	return h.hotelRepo.Create(createHotelRequest)
}

func (h *HotelService) Update(id domain.ID, hotel domain.CreateHotelRequest) error {
	return h.hotelRepo.Update(id, hotel)
}
