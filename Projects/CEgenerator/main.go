package main

import (
	"context"
	"log"
	"os"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cehttp "github.com/cloudevents/sdk-go/v2/protocol/http"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	typ := os.Getenv("TYPE")

	if typ == "" {
		typ = "type.fail"
	}

	path := "http://localhost:" + port + "/"
	log.Println("Sending Events at Path ", path)
	ctx := cloudevents.ContextWithTarget(context.Background(), path)

	p, err := cloudevents.NewHTTP()
	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}

	c, err := cloudevents.NewClient(p, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	for i := 0; ; i++ {
		e := cloudevents.NewEvent()
		e.SetType(typ)
		e.SetSource("https://aajhiudhadw.com/")
		_ = e.SetData(cloudevents.ApplicationJSON, map[string]interface{}{
			"id":      i,
			"message": "Hello Pablo!",
		})

		res := c.Send(ctx, e)
		if cloudevents.IsUndelivered(res) {
			log.Printf("Failed to send: %v", res)
		} else {
			var httpResult *cehttp.Result
			cloudevents.ResultAs(res, &httpResult)
			log.Printf("Sent %d with status code %d", i, httpResult.StatusCode)
		}

		time.Sleep(10000 * time.Millisecond)
	}
}
