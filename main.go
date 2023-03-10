package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler() {
	fmt.Println("function invoked!")
}

func main() {
	lambda.Start(Handler)
}
