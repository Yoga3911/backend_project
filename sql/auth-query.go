package sql

const Authentication = `SELECT * FROM users WHERE username = $1 AND password = $2`
