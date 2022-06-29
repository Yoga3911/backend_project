package sql

const Rollback = `CREATE OR REPLACE PROCEDURE rollbackMigration()
LANGUAGE plpgsql
AS $$
BEGIN
	DROP TABLE users;
	DROP TABLE role;
END;
$$;`

const CallRollback = `CALL rollbackMigration();`