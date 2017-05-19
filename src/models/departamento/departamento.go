package departamento

type Departamento struct {
	Id      int
	Nome    string
	Contato string
}

// func buscaTodos() []Departamento {
// 	rows, _ := db.CONEXAO.Query("select id, nome from Departamento")
// 	defer rows.Close()
// 	users := []Departamento{}
// 	for rows.Next() {
// 		u := Departamento{}
// 		rows.Scan(&u.Id, &u.Nome)
// 		users = append(users, u)
// 	}
// 	return users
// }

// func porId(id string) Departamento {
// 	u := Departamento{}
// 	err := db.CONEXAO.QueryRow("select id from Departamento where id = $1", id).Scan(&u.Id)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return u
// }
