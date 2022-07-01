package dto

type InsertProduct struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserId      string `json:"user_id" validate:"required"`
	CategoryId  int8   `json:"category_id" validate:"required"`
	CreatedAt   int64  `json:"created_at" validate:"required"`
	UpdatedAt   int64  `json:"updated_at" validate:"required"`
}
