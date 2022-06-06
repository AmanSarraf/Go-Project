package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", Index)

	fmt.Println("Server start listening on post: 8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello! %v ", html.EscapeString(r.URL.Path)) // r.URL.Path[1:] I don't know why this was used
}
