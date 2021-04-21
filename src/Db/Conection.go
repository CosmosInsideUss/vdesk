package Db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Conexao, _ = sql.Open("postgres", "user= password= dbname=vdesk sslmode=disable")
