package repositories

import "hotel/internal/domain"

type Hotel interface {
	GetById(id int64) (*domain.Hotel, error)
	GetAll(filter domain.GetHotelsRequest) ([]domain.Hotel, error)
	Create(createHotelRequest domain.CreateHotelRequest) (domain.ID, error)
	Update(id domain.ID, createHotelRequest domain.CreateHotelRequest) error
}
