package repository

import "github.com/jackc/pgx/v4/pgxpool"

type ProductR interface {
	GetAllProduct()
	GetProductById()
	InsertProduct()
	EditProduct()
	DeleteProduct()
}

type productR struct {
	db *pgxpool.Pool
}

func NewProductR() ProductR {
	return &productR{}
}

func (p *productR) GetAllProduct() {

}

func (p *productR) GetProductById() {

}

func (p *productR) InsertProduct() {

}

func (p *productR) EditProduct() {

}

func (p *productR) DeleteProduct() {

}
