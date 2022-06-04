package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BobrealnameStruct struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/readme", readdata)

	fmt.Println("listeninging on port 8080")
	http.ListenAndServe(":8080", nil)
}

// readdata expects a HTTP post request with a json body of `{"name":"bob"}`
func readdata(w http.ResponseWriter, req *http.Request) {
	bob := &BobStruct{}
	err := json.NewDecoder(req.Body).Decode(&bob)
	if err != nil {
		fmt.Println("error at readddata")
	}

	fmt.Println(bob.Name)

}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
