package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// book struct : MODEL
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Authon struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

//functions

func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {}

func updateBook(w http.ResponseWriter, r *http.Request) {}

func deleteBooks(w http.ResponseWriter, r *http.Request) {}

func main() {
	// init Router
	r := mux.NewRouter()

	// Mock data - @todo  - implement db
	books = append(books,
		Book{ID: "1", Isbn: "1101", Title: "Book One", Author: &Author{
			Firstname: "Arnab", Lastname: "Saha"}})
	books = append(books,
		Book{ID: "2", Isbn: "1102", Title: "Book Two", Author: &Author{
			Firstname: "Arpan", Lastname: "Saha"}})
	books = append(books,
		Book{ID: "3", Isbn: "1103", Title: "Book Three", Author: &Author{
			Firstname: "Nil", Lastname: "Saha"}})

	// create Route handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	// run server
	log.Fatal(http.ListenAndServe(":8000", r))
}
