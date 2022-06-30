package controllers

import (
	"crud/services"
	"crud/utils"

	"github.com/gofiber/fiber/v2"
)

type UserC interface {
	GetUserById(c *fiber.Ctx) error
}

type userC struct {
	userS services.UserS
}

func NewUserC(userS services.UserS) UserC {
	return &userC{
		userS: userS,
	}
}

func (u *userC) GetUserById(c *fiber.Ctx) error {
	user, err := u.userS.GetUserById(c.Context(), c.Params("id"))
	if err != nil {
		return utils.Response(c, 400, nil, err.Error(), false)
	}

	return utils.Response(c, 200, user, "Berhasil mendapatkan data user berdasarkan id", true)
}
