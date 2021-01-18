package handler

import "database/sql"

// DbStruct struct de conexão do banco, para o modelo singleton
type DbStruct struct {
	PostgreSQLString *sql.DB
}
