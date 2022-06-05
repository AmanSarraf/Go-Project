package main

import (
	"fmt"
	"net/http"
)

func main() {

	// we are telling main to check static directory
	fileserver := http.FileServer(http.Dir("./static"))

	//Handling root rout "/"
	http.Handle("/", fileserver) //<<<<<  " / " will erve as index.html
	//handler function for handling form.html
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server started at post :8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
		//or use log.Fatel(err) to get timestamp and logs
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	/*w is for response   -->what server sends back to the user
	  r is request        -->what user sends to the server*/

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported ", http.StatusNotFound)

	}

	fmt.Fprintf(w, "Hello ! success")

}
func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseform error %v", err)
		return
	}
	fmt.Fprintf(w, "Parseform Sucess! \n Getting data\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n Address = %s", name, address)

}
