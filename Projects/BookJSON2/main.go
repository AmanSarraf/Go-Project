package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func (b *BookHandler) book(w http.ResponseWriter, r *http.Request) {

	path := strings.Split(r.URL.Path, "/")
	if len(path) > 3 {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintln(w, "wrong path")
		return
	}

	switch r.Method {
	case http.MethodGet: //Retriving a book with books/{Id}

		{

			id, err := strconv.Atoi(path[2])
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Specify a valid Id")
				return
			}

			for _, value := range b.Books {
				if value.Id == id {
					fmt.Fprintln(w, "Book Found in record:\n")
					if err := json.NewEncoder(w).Encode(value); err != nil {
						log.Println(err)
					}
					return
				}

			}
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Book of Id %v doesn't exist\n", id)
			return

		}

	case http.MethodPost: // adding new book
		{
			if path[2] != "" {

				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Cannot specify any value while adding")
				return
			}

			newbook := Book{}
			_ = json.NewDecoder(r.Body).Decode(&newbook)
			for _, item := range b.Books {
				if newbook.Id == item.Id {

					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "Book with Id=%v already exists in record", newbook.Id)
					return
				}
			}
			b.Books = append(b.Books, newbook)
			fmt.Fprintln(w, "Added new Book")

			if err := json.NewEncoder(w).Encode(newbook); err != nil {
				log.Println(err)
			}
		}

	case http.MethodDelete: //Delete a specific book data
		{
			w.Header().Set("content-type", "application/json")

			id, err := strconv.Atoi(path[2])
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadGateway)
				fmt.Fprintln(w, "Specify a valid Id")
				return
			}

			for index, match := range b.Books {
				if match.Id == id {
					b.Books = append(b.Books[:index], b.Books[index+1:]...)
					fmt.Fprintln(w, "Deletion Successful")
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Book of Id %v doesn't exist\n", id)
			return
		}
	case http.MethodPut: //Update an existing book
		{
			if path[2] != "" {

				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Cannot specify any value while Updating")
				return
			}
			w.Header().Set("content-type", "application/json")
			jsonrequestfromweb := Book{}
			_ = json.NewDecoder(r.Body).Decode(&jsonrequestfromweb)

			for index, match := range b.Books {
				if match.Id == jsonrequestfromweb.Id {
					b.Books = append(b.Books[:index], b.Books[index+1:]...)
					b.Books = append(b.Books, jsonrequestfromweb)
					fmt.Fprintln(w, "Update successful")
					return

				}

			}
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Please Enter Valid Book Id in json body")

			return
		}
	}
}

func (b *BookHandler) allbooks(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet: //fetch the list of all books
		{

			if err := json.NewEncoder(w).Encode(b.Books); err != nil {
				log.Println(err)
			}

		}
	case http.MethodDelete: //delete all books
		{
			b.Books = nil
			fmt.Fprintln(w, "Library Cleared")
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
	http.HandleFunc("/books/", bs.book)
	// defines a route for fetching all books on get request
	//defines a route for deleting all books on DELETE request
	http.HandleFunc("/books", bs.allbooks)

	fmt.Println("server started at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
