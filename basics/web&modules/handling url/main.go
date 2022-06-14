package main

import (
	"fmt"
	"log"
	"net/url"
)

const myurl = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {

	//we don't to use myurl as stream , we need to PARSE it to something

	result, err := url.Parse(myurl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("type of myurl is %T", result)
}
