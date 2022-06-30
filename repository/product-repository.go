package repository

import (
	"crud/models"
	"crud/sql"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type ProductR interface {
	GetAllProduct(*fasthttp.RequestCtx, string) ([]*models.Product, error)
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

func (p *productR) GetAllProduct(ctx *fasthttp.RequestCtx, productId string) ([]*models.Product, error) {
	var product []*models.Product
	pgx, err := p.db.Query(ctx, sql.GetAllProduct)
	if err != nil {
		return product, err
	}

	for pgx.Next() {
		var p models.Product

		err = pgx.Scan(&p.Id, &p.Name, &p.Price, &p.Quantity, &p.Description, &p.UserId, &p.CategoryId, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			log.Println(err)
		}


		product = append(product, &p)
	}

	return product, nil
}

func (p *productR) GetProductById() {

}

func (p *productR) InsertProduct() {

}

func (p *productR) EditProduct() {

}

func (p *productR) DeleteProduct() {

}
