package services

import (
	"crud/dto"
	"crud/models"
	"crud/repository"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

type ProductS interface {
	GetAllProduct(*fasthttp.RequestCtx) ([]*models.Product, error)
	GetProductById(*fasthttp.RequestCtx, string) (models.Product, error)
	InsertProduct(*fasthttp.RequestCtx, dto.InsertProduct) (dto.InsertProduct, error)
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

func (p *productS) InsertProduct(ctx *fasthttp.RequestCtx, product dto.InsertProduct) (dto.InsertProduct, error) {
	err := p.productR.InsertProduct(ctx, product)
	if err != nil {
		return product, err
	}

	return product, nil
}
