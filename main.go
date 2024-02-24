package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type MyResponse struct {
	Message string `json:"Answer:"`
}

func handler(event MyEvent) (MyResponse, error) {
	log.Println("Hwllo world")
	return MyResponse{Message: fmt.Sprintf("%s is %d years old.", event.Name, event.Age)}, nil

}

func main() {
	lambda.Start(handler)
}
