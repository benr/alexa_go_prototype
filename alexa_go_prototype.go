package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// This "void" type acts as a holder to unmarshal the input JSON into
type Event interface{}

func HandleRequest(ctx context.Context, e Event) (string, error) {
	// Output the raw unindexed map[interface{}]interface{}:
	//log.Printf("Raw event: %v+\n", e)

	// Use Type Assertion to return an indexed map:
	myEvent := e.(map[string]interface{})

	// Iterate over Map Keys to examine input event:
	fmt.Println("---- Dumping Events Map: ----")
	for k, v := range myEvent {
		fmt.Printf("Key %v = %v\n", k, v)
	}
	fmt.Println("---- Done. ----")

	// Example of accessing map value via index:
	//fmt.Println("Key1 is ", myEvent["key1"])

	return fmt.Sprintf("Done!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
