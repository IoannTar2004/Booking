package domain

type Hotel struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Stars   int    `json:"stars"`
}
