package repository

import (
	"crud/models"
	"crud/sql"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type ProductR interface {
	GetAllProduct(*fasthttp.RequestCtx) ([]*models.Product, error)
	GetProductById()
	InsertProduct()
	EditProduct()
	DeleteProduct()
}

type productR struct {
	db *pgxpool.Pool
}

func NewProductR(db *pgxpool.Pool) ProductR {
	return &productR{
		db: db,
	}
}

func (p *productR) GetAllProduct(ctx *fasthttp.RequestCtx) ([]*models.Product, error) {
	var products []*models.Product
	pgx, err := p.db.Query(ctx, sql.GetAllProduct)
	if err != nil {
		return products, err
	}

	for pgx.Next() {
		var product models.Product

		err = pgx.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description, &product.UserId, &product.CategoryId, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			log.Println(err)
		}

		products = append(products, &product)
	}

	return products, nil
}

func (p *productR) GetProductById() {

}

func (p *productR) InsertProduct() {

}

func (p *productR) EditProduct() {

}

func (p *productR) DeleteProduct() {

}
