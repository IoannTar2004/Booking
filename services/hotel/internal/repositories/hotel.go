package repositories

import "hotel/internal/domain"

type Hotel interface {
	GetById(id int64) (*domain.Hotel, error)
}
