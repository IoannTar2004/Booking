package domain

type Hotel struct {
	Id      ID     `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Address string `json:"address"`
	Stars   int    `json:"stars"`
}

type GetHotelsRequest struct {
	Name   string `form:"name" binding:"max=64"`
	City   string `form:"city" binding:"max=32"`
	Stars  int    `form:"stars" binding:"lte=7"`
	Limit  int    `db:"limit" form:"limit" binding:"gte=1,lte=100"`
	Offset int    `db:"offset" form:"offset" binding:"gte=0"`
}

type CreateHotelRequest struct {
	Name    string `json:"name" binding:"required,max=64"`
	City    string `json:"city" binding:"required,max=32"`
	Address string `json:"address" binding:"required,max=100"`
	Stars   int    `json:"stars" binding:"required,gte=0,lte=7"`
}
