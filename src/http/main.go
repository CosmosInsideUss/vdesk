package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-pg/pg"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.PathPrefix("/js").Handler(http.StripPrefix("/js", http.FileServer(http.Dir("templates/js/"))))
	// //r.HandleFunc("/user/{id}", ViewHandler)
	fmt.Println(http.ListenAndServe(":8000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	//temp := template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))
	r.E
}
