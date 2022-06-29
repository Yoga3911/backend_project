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
	id SERIAL PRIMARY KEY,
	username VARCHAR(50) UNIQUE NOT NULL,
	email VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR(100),
	address VARCHAR(50),
	role_id SMALLINT NOT NULL,
	CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES role(id)
	);

END;
$$;`

const CallMigration = `CALL migration();`