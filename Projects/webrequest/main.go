package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://yourbasic.org/golang/http-server-example/"

func main() {
	fmt.Println("This is for simple http request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The type of response is %T\n", response)
	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)
}
