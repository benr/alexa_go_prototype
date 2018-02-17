package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event interface{}

func HandleRequest(ctx context.Context, e Event) (string, error) {
	//log.Printf("Got event: %v+\n", e)

	myEvent := e.(map[string]interface{})

	fmt.Println("---- Dumping Events Map: ----")
	for k, v := range myEvent {
		fmt.Printf("Key %v = %v\n", k, v)
	}
	fmt.Println("---- Done. ----")

	//fmt.Println("Key1 is ", myEvent["key1"])

	return fmt.Sprintf("Done!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
