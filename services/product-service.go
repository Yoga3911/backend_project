package services

import (
	"crud/dto"
	"crud/models"
	"crud/repository"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

type ProductS interface {
	GetAllProduct(*fasthttp.RequestCtx) ([]*models.Product, error)
	GetProductById(*fasthttp.RequestCtx, string) (models.Product, error)
	GetProductByCategoryId(*fasthttp.RequestCtx, dto.GetByCategory) ([]*models.Product, error)
	InsertProduct(*fasthttp.RequestCtx, dto.InsertProduct) (dto.InsertProduct, error)
	EditProduct(*fasthttp.RequestCtx, dto.EditProduct) (dto.EditProduct, error)
	DeleteProduct(*fasthttp.RequestCtx, dto.DeleteProduct) error
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

		err = data.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.UserId, &product.CategoryId, &product.IsDeleted, &product.CreatedAt, &product.UpdatedAt)
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
		Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.UserId, &product.CategoryId, &product.IsDeleted, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return product, errors.New("produk tidak ditemukan")
		}

		return product, err
	}

	return product, nil
}

func (p *productS) GetProductByCategoryId(ctx *fasthttp.RequestCtx, category dto.GetByCategory) ([]*models.Product, error) {
	var products []*models.Product

	data, err := p.productR.GetProductByCategoryId(ctx, category.Id)
	if err != nil {
		return products, err
	}

	for data.Next() {
		var product models.Product

		err = data.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.UserId, &product.CategoryId, &product.IsDeleted, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			log.Println(err)
		}

		products = append(products, &product)
	}

	return products, nil
}

func (p *productS) InsertProduct(ctx *fasthttp.RequestCtx, insertProduct dto.InsertProduct) (dto.InsertProduct, error) {
	insertProduct.Id = uuid.New().String()

	timeMili := time.Now().UnixMilli()
	insertProduct.CreatedAt = timeMili
	insertProduct.UpdatedAt = timeMili

	err := p.productR.InsertProduct(ctx, insertProduct)
	if err != nil {
		if strings.Contains(err.Error(), "fk_user_id") {
			return insertProduct, errors.New("invalid user")
		}

		return insertProduct, err
	}

	return insertProduct, nil
}

func (p *productS) EditProduct(ctx *fasthttp.RequestCtx, editProduct dto.EditProduct) (dto.EditProduct, error) {
	editProduct.UpdatedAt = time.Now().UnixMilli()

	err := p.productR.EditProduct(ctx, editProduct)
	if err != nil {
		return editProduct, err
	}

	return editProduct, nil
}

func (p *productS) DeleteProduct(ctx *fasthttp.RequestCtx, deleteProduct dto.DeleteProduct) error {
	deleteProduct.UpdatedAt = time.Now().UnixMilli()

	err := p.productR.DeleteProduct(ctx, deleteProduct)
	if err != nil {
		return err
	}

	return nil
}
