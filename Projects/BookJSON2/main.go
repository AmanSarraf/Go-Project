package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Name   string  `json:"bookname"`
	Id     int     `json:"id"`
	Price  int64   `json:"bookprice"`
	Isbn   string  `json:"isbn"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string
	Lastname  string
}

type Books struct {
	Books []Book `json:"book"`
}

//fake db
var books []Book

func (b *Books) book(w http.ResponseWriter, r *http.Request) {
	// if r.Method == POST
	switch r.Method {
	case "POST": //Retriving a book with Id number
		{
			w.Header().Set("content-type", "application/json")
			jsonreqestfromweb := &Book{}
			err := json.NewDecoder(r.Body).Decode(&jsonreqestfromweb)
			if err != nil {
				fmt.Println(err)
				return
			}

			json.NewEncoder(w).Encode(b.Books[jsonreqestfromweb.Id])
			break
		}
	case "PUT": // adding new book
		{

			w.Header().Set("content-type", "application/json")
			newbook := Book{}
			_ = json.NewDecoder(r.Body).Decode(&newbook)
			b.Books = append(b.Books, newbook)
			json.NewEncoder(w).Encode(newbook)
			break

		}
	case "DELETE": //Delete a specific book data
		{
			w.Header().Set("content-type", "application/json")
			jsonrequestfromweb := Book{}
			_ = json.NewDecoder(r.Body).Decode(&jsonrequestfromweb)
			b.Books = append(b.Books[:jsonrequestfromweb.Id], b.Books[jsonrequestfromweb.Id+1:]...)
			json.NewEncoder(w).Encode(b.Books)
			break
		}
	case "PATCH": //Update an existing book
		{
			w.Header().Set("content-type", "application/json")
			jsonrequestfromweb := Book{}
			_ = json.NewDecoder(r.Body).Decode(&jsonrequestfromweb)
			b.Books = append(b.Books[:jsonrequestfromweb.Id], b.Books[jsonrequestfromweb.Id+1:]...)
			b.Books = append(b.Books, jsonrequestfromweb)
			json.NewEncoder(w).Encode(b.Books)
			break
		}
	}
}

func (b *Books) allbooks(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET": //fetch the list of all books
		{
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(b.Books)
			break
		}
	case "DELETE": //delete all books
		{
			w.Header().Set("content-type", "application/json")
			b.Books = nil
			json.NewEncoder(w).Encode(b.Books)
			break

		}
	}
}

func main() {
	fmt.Println("Welcome")
	//seed value
	bs := &Books{}
	b1 := Book{Id: 0, Isbn: "24242242", Price: 856, Name: "Harry Potter", Author: &Author{Firstname: "JK", Lastname: "Roweling"}}
	b2 := Book{Id: 1, Isbn: "15454832", Price: 745, Name: "How to win friend", Author: &Author{Firstname: "Dan", Lastname: "Carnegi"}}
	bs.Books = append(bs.Books, b1)
	bs.Books = append(bs.Books, b2)

	// Defines a route for fetching a single book on GET request
	// Defines a route for delete a single book on DEELTE request
	// Defines a route for Posting a single book on POST request
	http.HandleFunc("/book", bs.book)
	// defines a route for fetching all books on get request
	//defines a route for deleting all books on DELETE request
	http.HandleFunc("/books", bs.allbooks)

	fmt.Println("server started at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
