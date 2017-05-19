package usuario

import (
	db "../../Db"
	departamento "../departamento"
)

type Usuario struct {
	Id              int
	Nome            string
	Email           string
	Senha           string
	Idade           int
	Departamento_id departamento.Departamento
}

func BuscaTodos() []Usuario {
	rows, err := db.Conexao.Query("select u.id, u.nome, u.email, u.idade, d.id, d.nome from usuario u join departamento d on u.departamento_id = d.id")
	if err != nil {

		panic(err)
	}
	defer rows.Close()
	users := []Usuario{}
	for rows.Next() {
		u := Usuario{}
		rows.Scan(&u.Id, &u.Nome, &u.Email, &u.Idade, &u.Departamento_id.Id, &u.Departamento_id.Nome)
		users = append(users, u)
	}
	return users

}

// func PorId(id string) Usuario {
// 	u := Usuario{}
// 	err := db.CONEXAO.QueryRow("select id from usuario where id = $1", id).Scan(&u.Id)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return u
// }

// func BuscaTodosJson() []Usuario {
// 	rows, _ := db.CONEXAO.Query("select * from usuario")
// 	defer rows.Close()
// 	users := []Usuario{}
// 	for rows.Next() {
// 		u := Usuario{}
// 		rows.Scan(&u.Nome)
// 		users = append(users, u)
// 	}
// 	return users
//}
