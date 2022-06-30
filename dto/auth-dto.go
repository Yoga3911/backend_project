package dto

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Register struct {
	Id       string `json:"id" `
	Username string `json:"username" validate:"required,min=6,max=50"`
	Email    string `json:"email" validate:"required,email,min=6,max=50"`
	Password string `json:"password" validate:"required,min=6,max=100"`
	Address  string `json:"address" validate:"required,min=6,max=50"`
}
