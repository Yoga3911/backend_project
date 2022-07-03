package dto

type GetByCategory struct {
	Id string `query:"categoryId"`
}

type InsertProduct struct {
	Id          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserId      string `json:"user_id" validate:"required"`
	CategoryId  int8   `json:"category_id" validate:"required"`
	IsDeleted   bool   `json:"is_deleted"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type EditProduct struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserId      string `json:"user_id" validate:"required"`
	CategoryId  int8   `json:"category_id" validate:"required"`
	UpdatedAt   int64  `json:"updated_at"`
}

type DeleteProduct struct {
	ProductId string `json:"product_id" validate:"required"`
	UserId    string `json:"user_id" validate:"required"`
	UpdatedAt int64  `json:"updated_at"`
}
