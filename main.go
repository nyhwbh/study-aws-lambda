package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (string, error) {
	myEnvVar := os.Getenv("APIkey")
	if myEnvVar == "" {
		fmt.Println("환경변수 APIkey가 설정되어 있지 않습니다.")
	} else {
		fmt.Printf("환경변수 APIkey의 값은 %s입니다.\n", myEnvVar)
	}
	return "AWS Lambda에서 환경변수를 가져왔습니다.", nil
}

func main() {
	lambda.Start(HandleRequest)
}
