package database

import (
	"database/sql"
	"errors"
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

	query := `
        SELECT h.id, h.name as name, c.name as city, h.address, h.stars FROM hotels h
	  	JOIN cities c ON h.city_id = c.id
        WHERE h.id = $1
    `
	err := h.db.Get(&hotel, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.HotelNotFoundError{}
	}
	if err != nil {
		return nil, err
	}

	return &hotel, nil
}

func (h *HotelRepository) GetAll(filter domain.GetHotelsRequest) ([]domain.Hotel, error) {
	query := `
        SELECT h.id, h.name as name, c.name as city, h.address, h.stars FROM hotels h
	  	JOIN cities c ON h.city_id = c.id
        WHERE (h.name ILIKE '%' || $1 || '%' OR $1 = '')
	    AND (c.name = $2 OR $2 = '')
        AND (h.stars = $3 OR $3 = 0)
        LIMIT $4 OFFSET $5
    `
	args := []interface{}{filter.Name, filter.City, filter.Stars, filter.Limit, filter.Offset}

	var hotels []domain.Hotel
	err := h.db.Select(&hotels, query, args...)
	if hotels == nil {
		return nil, domain.HotelNotFoundError{}
	}
	if err != nil {
		return nil, err
	}

	return hotels, nil
}
