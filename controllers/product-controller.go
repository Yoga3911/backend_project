package controllers

import (
	"crud/dto"
	"crud/services"
	"crud/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductC interface {
	GetAllProduct(*fiber.Ctx) error
	GetProductById(*fiber.Ctx) error
	InsertProduct(*fiber.Ctx) error
	EditProduct(*fiber.Ctx) error
	DeleteProduct(*fiber.Ctx) error
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

func (p *productC) InsertProduct(c *fiber.Ctx) error {
	var product dto.InsertProduct

	err := c.BodyParser(&product)
	if err != nil {
		return utils.Response(c, 400, nil, err.Error(), false)
	}

	if err := utils.StructValidator(product); err != nil {
		return utils.Response(c, 400, err, "There is something wrong!", false)
	}

	product, err = p.productS.InsertProduct(c.Context(), product)
	if err != nil {
		return utils.Response(c, 400, nil, err.Error(), false)
	}

	return utils.Response(c, 200, product, "Insert product success!", true)
}

func (p *productC) EditProduct(c *fiber.Ctx) error {
	var product dto.EditProduct

	err := c.BodyParser(&product)
	if err != nil {
		return utils.Response(c, 400, nil, err.Error(), false)
	}

	if err := utils.StructValidator(product); err != nil {
		return utils.Response(c, 400, err, "There is something wrong!", false)
	}

	product, err = p.productS.EditProduct(c.Context(), product)
	if err != nil {
		return utils.Response(c, 400, nil, err.Error(), false)
	}

	return utils.Response(c, 200, product, "Edit product success!", true)
}

func (p *productC) DeleteProduct(c *fiber.Ctx) error {
	var product dto.DeleteProduct

	err := c.BodyParser(&product)
	if err != nil {
		return utils.Response(c, 400, nil, err.Error(), false)
	}

	err = p.productS.DeleteProduct(c.Context(), product)
	if err != nil {
		return utils.Response(c, 400, nil, err.Error(), false)
	}

	return utils.Response(c, 200, product, "Delete product success!", true)
}
