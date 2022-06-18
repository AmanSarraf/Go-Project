package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
	"github.com/Azure/go-autorest/autorest"
)

func BatchReadFileRemoteImage(remoteImageURL string) {
	fmt.Println("-----------------------------------------")
	fmt.Println("BATCH READ FILE - remote")
	fmt.Println()
	endpointURL := os.Getenv("AZURE_VISION_URL")
	computerVisionKey := os.Getenv("AZURE_VISION_KEY")
	client := computervision.New(endpointURL)
	client.Authorizer = autorest.NewCognitiveServicesAuthorizer(computerVisionKey)
	computerVisionContext = context.Background()
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	textHeaders, err := client.BatchReadFile(computerVisionContext, remoteImage)
	if err != nil {
		log.Fatal(err)
	}

	// Use ExtractHeader from the autorest library to get the Operation-Location URL
	operationLocation := autorest.ExtractHeaderValue("Operation-Location", textHeaders.Response)

	numberOfCharsInOperationId := 36
	operationId := string(operationLocation[len(operationLocation)-numberOfCharsInOperationId : len(operationLocation)])
	// </snippet_read_call>

	// <snippet_read_response>
	readOperationResult, err := client.GetReadOperationResult(computerVisionContext, operationId)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the operation to complete.
	i := 0
	maxRetries := 10

	fmt.Println("Recognizing text in a remote image with the batch Read API ...")
	for readOperationResult.Status != computervision.Failed &&
		readOperationResult.Status != computervision.Succeeded {
		if i >= maxRetries {
			break
		}
		i++

		fmt.Printf("Server status: %v, waiting %v seconds...\n", readOperationResult.Status, i)
		time.Sleep(1 * time.Second)

		readOperationResult, err = client.GetReadOperationResult(computerVisionContext, operationId)
		if err != nil {
			log.Fatal(err)
		}
	}
	// </snippet_read_response>

	// <snippet_read_display>
	// Display the results.
	fmt.Println()
	for _, recResult := range *(readOperationResult.RecognitionResults) {
		for _, line := range *recResult.Lines {
			fmt.Println(*line.Text)
		}
	}
	// </snippet_read_display>
	fmt.Println()
}
