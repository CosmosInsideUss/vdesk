package main

import (
	//"../Db"
	// _ "../models/sla"
	// ticket "../models/ticket"
	// _ "../models/usuario"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-pg/pg"
	// "github.com/gorilla/mux"
	_ "github.com/lib/pq"
	// "net/http"
	// "text/template"
)

type Departamento struct {
	Id      int
	Nome    string
	Contato string
}
type Usuario struct {
	Id              int
	Nome            string
	Email           string
	Senha           string
	Idade           int
	Departamento_id Departamento
}

func main() {
	conexao, err := sql.Open("postgres", "user=postgres password=12345 dbname=vdesk sslmode=disable")
	if err != nil {
		panic(err)
	}
	rows, err := conexao.Query("select u.id, u.nome, u.email, u.senha, u.idade, d.id from usuario u join departamento d on u.departamento_id = d.id")
	if err != nil {
		panic(err)
	}
	usuarios := []Usuario{}
	for rows.Next() {
		u := Usuario{}
		rows.Scan(&u.Id, &u.Nome, &u.Email, &u.Senha, &u.Idade, &u.Departamento_id.Id)
		usuarios = append(usuarios, u)
	}

	lista, _ := json.Marshal(usuarios)

	fmt.Println(string(lista))

	// r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	// //r.HandleFunc("/sex", UsuariosJson)
	// //r.HandleFunc("/ticketsPendentes", TicketsIndex)

	// r.PathPrefix("/js").Handler(http.StripPrefix("/js", http.FileServer(http.Dir("templates/js/"))))
	// //r.HandleFunc("/user/{id}", ViewHandler)
	// fmt.Println(http.ListenAndServe(":8000", r))
}

// func HomeHandler(w http.ResponseWriter, r *http.Request) {

// 	temp := template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))
// 	temp.ExecuteTemplate(w, "layout.html", ticket.BuscaTodos())

// }

// func ViewHandler(w http.ResponseWriter, r *http.Request) {
// 	id := mux.Vars(r)["id"]
// 	temp := template.Must(template.ParseFiles("templates/user.html"))
// 	temp.ExecuteTemplate(w, "user.html", usuario.PorId(id))

// }

// func UsuariosJson(w http.ResponseWriter, h *http.Request) {
// 	v := usuario.BuscaTodos()
// 	js, err := json.Marshal(v)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(js)
// }

// func TicketsIndex(w http.ResponseWriter, h *http.Request) {
// 	t := ticket.BuscaTodos()
// 	js, err := json.Marshal(t)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(js)
//}
