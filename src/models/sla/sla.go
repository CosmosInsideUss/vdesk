package sla

import (
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "vdesk"
)

type Sla struct {
	Id      int
	Valor   string
	Usuario Usuario
}
type Usuario struct {
	Id    int
	Nome  string
	Email string
	Senha string
	Idade int
}

func SearchAll() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	
	rows, err := db.Query("select * from slas")
	if err != nil {

		panic(err)
	}

	slas := []Sla{}
	for rows.Next() {
		r := Sla{}
		rows.Scan(&r.Id, &r.Valor, &r.Usuario)
		slas = append(slas, r)
		fmt.Println(r.Usuario.Id)

	}

	// sla1 := &Sla{
	// 	Valor: "10min",
	// }
	// err := db.Conexao.Insert(sla1)
	// if err != nil {
	// 	panic(err)
	// }

	// Select user by primary key.
	// sla := Sla{Id: 2}
	// err := db.Conexao.Select(&sla)
	// if err != nil {
	// 	panic(err)
	// }

	//Select all users.

}
