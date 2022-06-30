package models

type User struct {
	Id        string `json:"id" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Address   string `json:"address" validate:"required"`
	RoleId    int8   `json:"role_id" validate:"required"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Token     string `json:"token"`
}
