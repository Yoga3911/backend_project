package sql

const Authentication = `SELECT * FROM users WHERE username = $1 AND password = $2`

const InsertUser = `INSERT INTO users (username, email, password, address) 
					VALUES ($1, $2, $3, $4)`
