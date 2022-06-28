package controllers

import (
	"crud/dto"
	"crud/helpers"
	"crud/services"

	"github.com/gofiber/fiber/v2"
)

type AuthC interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type authC struct {
	authS services.AuthS
}

func NewAuthC(auths services.AuthS) AuthC {
	return &authC{
		authS: auths,
	}
}

func (a *authC) Login(c *fiber.Ctx) error {
	var loginData dto.Login
	c.BodyParser(&loginData)
	if err := helpers.EmptyChecker(loginData); err != nil {
		return helpers.Response(c, 400, err, "Login gagal!", false)
	}

	user, err := a.authS.LoginUser(loginData, c.Context())
	if err != nil {
		return helpers.Response(c, 400, nil, err.Error(), false)
	}

	return helpers.Response(c, 200, user, "Login berhasil!", true)
}

func (a *authC) Register(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Register berhasil!",
		"data":    nil,
	})
}
