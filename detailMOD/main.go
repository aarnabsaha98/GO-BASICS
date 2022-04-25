package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello!  Mod in GoLang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	//running server
	// http.ListenAndServe(":4000", r)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("hey !")
}

func serverHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> Welcome to GoLang </h1>"))
}
