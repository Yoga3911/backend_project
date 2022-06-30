package sql

const Authentication = `SELECT * FROM users WHERE username = $1`

const InsertUser = `INSERT INTO users (id, username, email, password, address, role_id, created_at, updated_at) 
					VALUES ($1, $2, $3, $4, $5, 1, $6, $7)`
