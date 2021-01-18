package psql

import (
	"database/sql"
	"log"
	"sync"

	gconcat "github.com/ancogamer/product5/api.crud/pkg/gconcat"
	_ "github.com/lib/pq"
)

var (
	once    sync.Once
	err     error
	dbLocal *sql.DB
)

// Connect faz 1 conexão com o banco
func Connect() *sql.DB {
	once.Do(func() {
		if dbLocal != nil {
			return
		}
		var dbInfo string
		dbInfo = gconcat.Concat("host=", dbHost, " port=",
			dbPort, " user=", dbUser,
			" password=", dbPassword, " dbname=", dbName)

		if dbLocal, err = sql.Open(dbSorce, dbInfo); err != nil {
			log.Println("error open:", err)
			return
		}
	})
	return dbLocal
}
