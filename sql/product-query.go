package sql

const GetAllProduct = `SELECT * FROM products WHERE is_deleted = false`

const GetProductById = `SELECT * FROM products WHERE id = $1`

const InsertProduct = `INSERT INTO products (id, name, price, quantity, description, user_id, category_id, is_deleted, created_at, updated_at) 
						VALUES ($1, $2, $3, $4, $5, $6, $7, false, $8, $9)`

const EditProduct = `UPDATE products SET 
					name = $3, price = $4, quantity = $5, description = $6, category_id = $7, updated_at = $8 
					WHERE id = $1 AND user_id = $2`

const DeleteProduct = `UPDATE products SET is_deleted = true, updated_at = $3 WHERE id = $1 AND user_id = $2`
