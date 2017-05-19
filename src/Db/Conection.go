package Db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Conexao, _ = sql.Open("postgres", "user=postgres password=12345 dbname=vdesk sslmode=disable")
