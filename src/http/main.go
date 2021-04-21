package main

import (
	//db "../Db"

	//"database/sql"

	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-pg/pg"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	usuario "github.com/pnkfd/vdesk/src/models"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	// //r.HandleFunc("/user/{id}", ViewHandler)
	fmt.Println(http.ListenAndServe(":8000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("../templates/temp.html"))
	User := usuario.Usuario{1, "maicon"}
	temp.Execute(w, User)

}
