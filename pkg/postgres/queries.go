package postgres

const (
	CREATE_DB           = `CREATE DATABASE "%s"`
	ALTER_DB_OWNER      = `ALTER DATABASE "%s" OWNER TO "%s"`
	CREATE_GROUP_ROLE   = `CREATE ROLE "%s"`
	CREATE_USER_ROLE    = `CREATE ROLE "%s" WITH LOGIN PASSWORD '%s'`
	GRANT_ROLE          = `GRANT "%s" TO "%s"`
	ALTER_USER_SET_ROLE = `ALTER USER "%s" SET ROLE "%s"`
	REVOKE_ROLE         = `REVOKE "%s" FROM "%s"`
	UPDATE_PASSWORD     = `ALTER ROLE "%s" WITH PASSWORD '%s'`
	DROP_ROLE           = `DROP ROLE "%s"`
	DROP_DATABASE       = `DROP DATABASE "%s"`
	REASIGN_OBJECTS     = `REASSIGN OWNED BY "%s" TO "%s"`
)
