package main

import (
	"fmt"
	"net/http"
)

const url string = "https://www.youtube.com/watch?v=ru53LpdVHn4&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=25"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	//databytes, err := ioutil.ReadAll(response.Body)

	//
	//content := string(databytes)
	// fmt.Println(content)

}
