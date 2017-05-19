package main

import (
	//db "../Db"
	ticket "../models/ticket"
	usuario "../models/usuario"
	//"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-pg/pg"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"text/template"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/ticketsPendentes", TicketsPendentes)
	r.HandleFunc("/ticketsUsuario", TicketsUsuario)
	r.PathPrefix("/js").Handler(http.StripPrefix("/js", http.FileServer(http.Dir("templates/js/"))))
	// //r.HandleFunc("/user/{id}", ViewHandler)
	fmt.Println(http.ListenAndServe(":8000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))
	temp.ExecuteTemplate(w, "layout.html", ticket.Pendentes())

}

func UsuariosJson(w http.ResponseWriter, h *http.Request) {
	t := usuario.BuscaTodos()
	js, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// func ViewHandler(w http.ResponseWriter, r *http.Request) {
// 	id := mux.Vars(r)["id"]
// 	temp := template.Must(template.ParseFiles("templates/user.html"))
// 	temp.ExecuteTemplate(w, "user.html", usuario.PorId(id))

// }

func TicketsPendentes(w http.ResponseWriter, h *http.Request) {
	v := ticket.Pendentes()
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func TicketsUsuario(w http.ResponseWriter, h *http.Request) {
	v := ticket.PorUsuario()
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
