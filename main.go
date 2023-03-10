package main

import (
	"context"
	"fmt"
	"studylambda/bringdata"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (string, error) {
	data := bringdata.CheckNewData()

	fmt.Println(data)

	return "AWS Lambda에서 환경변수를 가져왔습니다.", nil
}

func main() {
	lambda.Start(HandleRequest)
}
