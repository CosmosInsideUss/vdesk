package empresa

type Empresa struct {
	Id      int
	Nome    string
	Contato string
}

// func buscaTodos() []Empresa {
// 	rows, _ := db.CONEXAO.Query("select id, nome from Empresa")
// 	defer rows.Close()
// 	users := []Empresa{}
// 	for rows.Next() {
// 		u := Empresa{}
// 		rows.Scan(&u.Id, &u.Nome)
// 		users = append(users, u)
// 	}
// 	return users
// }

// func porId(id string) Empresa {
// 	u := Empresa{}
// 	err := db.CONEXAO.QueryRow("select id from Empresa where id = $1", id).Scan(&u.Id)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return u
// }
