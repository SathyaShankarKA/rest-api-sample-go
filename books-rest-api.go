package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

var books []Book

//Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get a book by id
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//Post a book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)
	newBook.ID = strconv.Itoa(len(books) + 1)
	books = append(books, newBook)
	json.NewEncoder(w).Encode(newBook)
}

//Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
}

//Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
}

func main() {
	//generate sample data
	books = append(books, Book{ID: "1", Name: "Five Point Someone", Author: "Chetan Bhagat"}, Book{ID: "2", Name: "Two States", Author: "Chetan Bhagat"})

	fmt.Println("Initialized sample data")

	//initialize Router
	router := mux.NewRouter()

	//rest endpoints
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))

}
