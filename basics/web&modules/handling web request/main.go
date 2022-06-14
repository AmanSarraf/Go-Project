package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://github.com/AmanSarraf/Go-Project/blob/main/Projects/BookJSON/main.go"

func main() {

	fmt.Println("Simple URL request reader")

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
