package controllers

import (
	"crud/services"

	"github.com/gofiber/fiber/v2"
)

type AuthC interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type authC struct {
	authS *services.AuthS
}

func NewAuthC(auths services.AuthS) AuthC {
	return &authC{
		authS: &auths,
	}
}

func (a *authC) Login(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Login berhasil!",
		"data":    nil,
	})
}

func (a *authC) Register(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Register berhasil!",
		"data":    nil,
	})
}
