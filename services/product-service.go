package services

import (
	"crud/models"
	"crud/repository"

	"github.com/valyala/fasthttp"
)

type ProductS interface {
	GetAllProduct(*fasthttp.RequestCtx) ([]*models.Product, error)
}

type productS struct {
	productR repository.ProductR
}

func NewProductR(productR repository.ProductR) ProductS {
	return &productS{
		productR: productR,
	}
}

func (p *productS) GetAllProduct(ctx *fasthttp.RequestCtx) ([]*models.Product, error) {
	products, err := p.productR.GetAllProduct(ctx)
	if err != nil {
		return products, err
	}

	return products, nil
}
