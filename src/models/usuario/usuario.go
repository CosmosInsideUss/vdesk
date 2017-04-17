package usuario

import (
	_ "../../Db"
	_ "fmt"
)

type Usuario struct {
	Id           int
	Nome         string
	Email        string
	Senha        string
	Idade        int
	Departamento string
}

// func BuscaTodos() []Usuario {
// 	rows, _ := db.CONEXAO.Query("select u.id, u.nome, u.email, u.idade, u.senha, d.nome " +
// 		"as departamento from usuario u join departamento d on u.departamento_id = d.id ORDER BY u.id ASC")
// 	// if err != nil {

// 	// 	panic(err)
// 	// }
// 	defer rows.Close()
// 	users := []Usuario{}
// 	for rows.Next() {
// 		u := Usuario{}
// 		if err := rows.Scan(&u.Id, &u.Nome, &u.Email, &u.Idade, &u.Senha, &u.Departamento); err != nil {
// 			panic(err)
// 		}
// 		users = append(users, u)
// 	}
// 	return users

// }

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
