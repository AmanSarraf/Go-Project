package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func main() {

	ctx := context.Background()
	p, err := cloudevents.NewHTTP()
	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}

	h, err := cloudevents.NewHTTPReceiveHandler(ctx, p, receive)
	if err != nil {
		log.Fatalf("failed to create handler: %s", err.Error())
	}

	http.Handle("/", h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/test", test)

	log.Printf("server starts at port: %v", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is test")
}

func receive(ctx context.Context, event cloudevents.Event) {
	if event.Type() != "type.fail" {
		fmt.Printf("Got an Event: %v", event)
	}

}
