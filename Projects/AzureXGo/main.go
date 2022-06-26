package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

var computerVisionContext context.Context

type myurl struct {
	Url string `json:"url"`
}

func main() {

	port := os.Getenv("PORT")
	http.HandleFunc("/", compvs)
	if port == "" {
		port = "8080"
	}

	fmt.Println("server started at port " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Panicln(err)
	}

}
