package models

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
	UserId      string `json:"user_id"`
	CategoryId  int8   `json:"category_id"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
