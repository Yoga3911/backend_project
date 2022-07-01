package controllers

import (
	"crud/services"
	"crud/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductC interface {
	GetAllProduct(*fiber.Ctx) error
	GetProductById(*fiber.Ctx) error
}

type productC struct {
	productS services.ProductS
}

func NewProductC(productS services.ProductS) ProductC {
	return &productC{
		productS: productS,
	}
}

func (p *productC) GetAllProduct(c *fiber.Ctx) error {
	products, err := p.productS.GetAllProduct(c.Context())
	if err != nil {
		return utils.Response(c, 400, nil, "Get All product data failed!", false)
	}

	return utils.Response(c, 200, products, "Get All product data success!", true)
}

func (p *productC) GetProductById(c *fiber.Ctx) error {
	products, err := p.productS.GetProductById(c.Context(), c.Params("productId"))
	if err != nil {
		return utils.Response(c, 400, nil, err.Error(), false)
	}

	return utils.Response(c, 200, products, "Get product data success!", true)
}
