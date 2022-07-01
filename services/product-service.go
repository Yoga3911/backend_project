package services

import (
	"crud/models"
	"crud/repository"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

type ProductS interface {
	GetAllProduct(*fasthttp.RequestCtx) ([]*models.Product, error)
	GetProductById(*fasthttp.RequestCtx, string) (models.Product, error)
}

type productS struct {
	productR repository.ProductR
}

func NewProductS(productR repository.ProductR) ProductS {
	return &productS{
		productR: productR,
	}
}

func (p *productS) GetAllProduct(ctx *fasthttp.RequestCtx) ([]*models.Product, error) {
	var products []*models.Product

	data, err := p.productR.GetAllProduct(ctx)
	if err != nil {
		return products, err
	}

	for data.Next() {
		var product models.Product

		err = data.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.UserId, &product.CategoryId, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			log.Println(err)
		}

		products = append(products, &product)
	}

	return products, nil
}

func (p *productS) GetProductById(ctx *fasthttp.RequestCtx, productId string) (models.Product, error) {
	var product models.Product

	err := p.productR.GetProductById(ctx, productId).
		Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.UserId, &product.CategoryId, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return product, fmt.Errorf("Produk tidak ditemukan!")
		}

		return product, err
	}

	return product, nil
}
