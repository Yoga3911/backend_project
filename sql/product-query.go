package sql

const GetAllProduct = `SELECT * FROM products`

const GetProductById = `SELECT * FROM products WHERE id = $1`

const InsertProduct = `INSERT INTO products (id, name, price, quantity, description, user_id, category_id, created_at, updated_at) 
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

const EditProduct = `UPDATE products SET 
					name = $2, price $3, quantity = $4, description = $5, category_id = $6, updated_at = $7 
					WHERE id = $1`

const DeleteProduct = `DELETE FROM products WHERE id = $1`
