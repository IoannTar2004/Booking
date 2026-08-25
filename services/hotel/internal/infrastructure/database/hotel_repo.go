package database

import (
	"database/sql"
	"errors"
	"fmt"
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

func (h *HotelRepository) Create(hotel domain.CreateHotelRequest) (domain.ID, error) {
	query := `
		INSERT INTO hotels (name, city_id, address, stars) 
		VALUES ($1, (SELECT id FROM cities WHERE cities.name = $2), $3, $4)
		RETURNING id
		`

	var hotelId domain.ID
	args := []interface{}{hotel.Name, hotel.City, hotel.Address, hotel.Stars}
	err := h.db.QueryRowx(query, args...).Scan(&hotelId)
	if err != nil {
		return 0, err
	}

	return hotelId, nil
}

func (h *HotelRepository) Update(id domain.ID, hotel domain.CreateHotelRequest) error {
	query := `
		UPDATE hotels SET name = $1, city_id = (SELECT id FROM cities WHERE cities.name = $2),
	    address = $3, stars = $4 WHERE id = $5
	`
	args := []interface{}{hotel.Name, hotel.City, hotel.Address, hotel.Stars, id}
	rows, err := h.db.Exec(query, args...)
	fmt.Println(rows)

	return err
}