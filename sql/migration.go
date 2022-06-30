package sql

const Migration = `CREATE OR REPLACE PROCEDURE migration()
LANGUAGE plpgsql
AS $$
BEGIN
	CREATE TABLE role (
	id SERIAL PRIMARY KEY,
	role VARCHAR(10) NOT NULL
	);
	
	INSERT INTO role(role) VALUES('Customer');
	INSERT INTO role(role) VALUES('Seller');
	INSERT INTO role(role) VALUES('Admin');
	
	CREATE TABLE users (
	id VARCHAR(50) PRIMARY KEY,
	username VARCHAR(50) UNIQUE NOT NULL,
	email VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR(100),
	address VARCHAR(50),
	role_id SMALLINT NOT NULL,
	created_at BIGINT NOT NULL,
	updated_at BIGINT NOT NULL,
	CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES role(id)
	);
	
	CREATE TABLE category (
		id SERIAL PRIMARY KEY,
		category VARCHAR(50) NOT NULL
	);

	INSERT INTO category(category) VALUES('Laptop');
	INSERT INTO category(category) VALUES('TV');
	INSERT INTO category(category) VALUES('HP');
	INSERT INTO category(category) VALUES('Kamera');
	INSERT INTO category(category) VALUES('Baju');
	INSERT INTO category(category) VALUES('Celana');
	INSERT INTO category(category) VALUES('Topi');
	INSERT INTO category(category) VALUES('Kacamata');
	INSERT INTO category(category) VALUES('Tas');

	CREATE TABLE products (
		id VARCHAR(50) PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		price INT NOT NULL,
		quantity INT NOT NULL,
		description TEXT NOT NULL,
		user_id VARCHAR(50) NOT NULL,
		category_id INT NOT NULL,
		created_at BIGINT NOT NULL,
		updated_at BIGINT NOT NULL,
		CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id),
		CONSTRAINT fk_category_id FOREIGN KEY (category_id) REFERENCES category(id)
	);
END;
$$;`

const CallMigration = `CALL migration();`
