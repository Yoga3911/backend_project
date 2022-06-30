package dto

type UserLogin struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	RoleId    int8   `json:"role_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Token     string `json:"token"`
}