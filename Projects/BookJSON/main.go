package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// BOOK STRUCT (model)

type book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//*******Functions************

//Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//add new book
func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var newbook book
	_ = json.NewDecoder(r.Body).Decode(&newbook)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	newbook.ID = strconv.Itoa(r1.Intn(100))
	books = append(books, newbook)
	json.NewEncoder(w).Encode(newbook)

}

//update book
func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")

	params := mux.Vars(r)
	for index, item := range books {

		if item.ID == params["ID"] {
			books = append(books[:index], books[index+1:]...) // <----

			var newbook book
			_ = json.NewDecoder(r.Body).Decode(&newbook)

			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			newbook.ID = strconv.Itoa(r1.Intn(100))
			books = append(books, newbook)
			json.NewEncoder(w).Encode(newbook)
			return

		}

	}

	json.NewEncoder(w).Encode(books)
}

//Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "Application/json")

	params := mux.Vars(r)
	for index, item := range books {

		if item.ID == params["ID"] {
			books = append(books[:index], books[index+1:]...) // <---- will delete
			break
		}

	}
	json.NewEncoder(w).Encode(books)

}

//Get a single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r) // get the params

	for _, value := range books {

		if value.ID == params["ID"] {
			json.NewEncoder(w).Encode(value)
			return
		}

	}
	json.NewEncoder(w).Encode(&book{})

}

// slice for holding book data
var books []book

func main() {

	fmt.Print("*********************WELCOME TO BOOK JSON PROJECT*************************\n")
	Router := mux.NewRouter()

	//TEST DATA - todo implement db

	books = append(books, book{ID: "1", Isbn: "24242242", Title: "Harry Potter", Author: &Author{Firstname: "JK", Lastname: "Roweling"}})
	books = append(books, book{ID: "2", Isbn: "15454832", Title: "How to win friend", Author: &Author{Firstname: "Dan", Lastname: "Carnegi"}})

	Router.HandleFunc("/books", getBooks).Methods("GET")
	Router.HandleFunc("/books", createBooks).Methods("POST")
	Router.HandleFunc("/books{ID}", updateBooks).Methods("PUT")
	Router.HandleFunc("/books/{ID}", deleteBook).Methods("DELETE")
	Router.HandleFunc("/books/{ID}", getBook).Methods("GET")
	fmt.Println("server started at port 8080")
	if err := http.ListenAndServe(":8080", Router); err != nil {
		log.Fatal(err)
	}

}
