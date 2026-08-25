package domain

type Hotel struct {
	Id      int64  `json:"id"`
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
