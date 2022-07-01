package repository

import (
	"crud/sql"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type ProductR interface {
	GetAllProduct(*fasthttp.RequestCtx) (pgx.Rows, error)
	GetProductById(*fasthttp.RequestCtx, string) pgx.Row
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

func (p *productR) GetAllProduct(ctx *fasthttp.RequestCtx) (pgx.Rows, error) {
	return p.db.Query(ctx, sql.GetAllProduct)
}

func (p *productR) GetProductById(ctx *fasthttp.RequestCtx, productId string) pgx.Row {
	return p.db.QueryRow(ctx, sql.GetProductById, productId)
}

func (p *productR) InsertProduct() {

}

func (p *productR) EditProduct() {

}

func (p *productR) DeleteProduct() {

}
