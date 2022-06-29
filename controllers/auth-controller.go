package controllers

import (
	"crud/dto"
	"crud/services"
	"crud/utils"

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

	if err := utils.StructValidator(loginData); err != nil {
		return utils.Response(c, 400, err, "Anda yang salah!", false)
	}

	user, err := a.authS.LoginUser(loginData, c.Context())
	if err != nil {
		return utils.Response(c, 400, nil, err.Error(), false)
	}

	return utils.Response(c, 200, user, "Login berhasil!", true)
}

func (a *authC) Register(c *fiber.Ctx) error {
	var registerData dto.Register
	c.BodyParser(&registerData)

	if err := utils.StructValidator(registerData); err != nil {
		return utils.Response(c, 400, err, "Anda yang salah!", false)
	}

	err := utils.InputChecker(registerData.Email, registerData.Password, registerData.Username, registerData.Address)
	if err != nil {
		return utils.Response(c, 400, nil, err.Error(), false)

	}

	err = a.authS.RegisterUser(registerData, c.Context())
	if err != nil {
		return utils.Response(c, 400, nil, err.Error(), false)
	}

	return utils.Response(c, 200, nil, "Register berhasil!", true)
}
