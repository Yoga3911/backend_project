package models

type User struct {
	Id       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address" validate:"required"`
	RoleId   int8   `json:"role_id" validate:"required"`
}
