package database

import (
	"hotel/internal/domain"
	"hotel/internal/repositories"

	"github.com/jmoiron/sqlx"
)

type HotelRepository struct {
	db *sqlx.DB
}

func NewHotelRepository(db *sqlx.DB) repositories.Hotel {
	return &HotelRepository{db: db}
}

func (h *HotelRepository) GetById(id int64) (*domain.Hotel, error) {
	var hotel domain.Hotel
	query := "SELECT * FROM hotels WHERE id=$1"
	err := h.db.Get(&hotel, query, id)
	if err != nil {
		return nil, err
	}

	return &hotel, nil
}
