package sql

const Authentication = `SELECT * FROM users WHERE username = $1 AND password = $2`

const InsertUser = `INSERT INTO users (username, email, password, address, role_id) 
					VALUES ($1, $2, $3, $4, 1)`
