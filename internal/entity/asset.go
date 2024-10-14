package entity

type Asset struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Owner_id    int64   `json:"owner_id"`
}
