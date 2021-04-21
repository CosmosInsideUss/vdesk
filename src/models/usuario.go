package models

import (
	banco "github.com/pnkfd/vdesk/src/banco"
)

type Usuario struct {
	Id   int
	Nome string
}

func BuscaTodos() []Usuario {
	rows, err := banco.Conexao.Query("select u.id, u.nome from usuario")
	if err != nil {

		panic(err)
	}
	defer rows.Close()
	users := []Usuario{}
	for rows.Next() {
		u := Usuario{}
		rows.Scan(&u.Id, &u.Nome)
		users = append(users, u)
	}
	return users

}
