package domain

type HotelNotFoundError struct{}

func (h HotelNotFoundError) Error() string {
	return "Hotel Not Found"
}
