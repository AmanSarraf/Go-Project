package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

type BookHandler struct {
	Books []Book `json:"book"`
}

//fake db
var books []Book

func (b *BookHandler) book(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: //Retriving a book with Id number
		{

			val := r.FormValue("id")
			if val == "" {
				if err := json.NewEncoder(w).Encode(b.Books); err != nil {
					log.Println(err)
					return
				}
				return
			}

			if val == "break" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			id, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}

			defer r.Body.Close()
			for i, value := range b.Books {
				if value.Id == id {
					if err := json.NewEncoder(w).Encode(b.Books[i]); err != nil {
						log.Println(err)

					}
					return
				}

			}

		}

	case http.MethodPost: // adding new book
		{
			newbook := Book{}
			_ = json.NewDecoder(r.Body).Decode(&newbook)
			for _, item := range b.Books {
				if newbook.Id == item.Id {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}
			b.Books = append(b.Books, newbook)

			if err := json.NewEncoder(w).Encode(newbook); err != nil {
				log.Println(err)
			}
		}

	case http.MethodDelete: //Delete a specific book data
		{
			w.Header().Set("content-type", "application/json")
			jsonrequestfromweb := Book{}
			_ = json.NewDecoder(r.Body).Decode(&jsonrequestfromweb)
			b.Books = append(b.Books[:jsonrequestfromweb.Id], b.Books[jsonrequestfromweb.Id+1:]...)
			//json.NewEncoder(w).Encode(b.Books)

		}
	case http.MethodPut: //Update an existing book
		{
			w.Header().Set("content-type", "application/json")
			jsonrequestfromweb := Book{}
			_ = json.NewDecoder(r.Body).Decode(&jsonrequestfromweb)
			b.Books = append(b.Books[:jsonrequestfromweb.Id], b.Books[jsonrequestfromweb.Id+1:]...)
			b.Books = append(b.Books, jsonrequestfromweb)
			//json.NewEncoder(w).Encode(b.Books)

		}
	}
}

func (b *BookHandler) allbooks(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet: //fetch the list of all books
		{
			//w.Header().Set("content-type", "application/json")
			if err := json.NewEncoder(w).Encode(b.Books); err != nil {
				log.Println(err)
			}

		}
	case http.MethodDelete: //delete all books
		{
			w.Header().Set("content-type", "application/json")
			b.Books = nil
			//json.NewEncoder(w).Encode(b.Books)

		}
	}
}

func main() {
	fmt.Println("Welcome")

	//seed value
	bs := &BookHandler{}
	b1 := Book{Id: 23, Isbn: "24242242", Price: 856, Name: "Harry Potter", Author: &Author{Firstname: "JK", Lastname: "Roweling"}}
	b2 := Book{Id: 12, Isbn: "15454832", Price: 745, Name: "How to win friend", Author: &Author{Firstname: "Dan", Lastname: "Carnegi"}}
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
