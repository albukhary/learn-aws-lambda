package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

func Handler(event Inputevent) (string, error) {
	fmt.Println("Assalamu alaykum ",event.FirstName, event.LastName)

	return "It worked", nil
}

type Inputevent struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

