package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
)

func main() {
	staticC := controllers.NewStatic()
	userC := controllers.NewUsers()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/signup", userC.New).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}
