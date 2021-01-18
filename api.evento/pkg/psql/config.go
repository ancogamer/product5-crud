package psql

import (
	"os"
)

var (
	dbName     = os.Getenv("DB_NAME")
	dbHost     = os.Getenv("DB_HOST")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbPort     = "5432"
	dbSSL      = "verify-ca"
	dbSorce    = "postgres"
)
