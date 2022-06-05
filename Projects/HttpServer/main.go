//HTTP Server
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", Helloserver)
	http.ListenAndServe(":3000", nil)

}

func Helloserver(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Namaste! \n")
	io.WriteString(w, "Jai Mahakal! \n")

}
