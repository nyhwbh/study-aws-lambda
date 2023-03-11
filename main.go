package main

import (
	"context"
	"studylambda/savedata"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (string, error) {
	savedata.SaveDataAll()

	return "AWS Lambda에서 환경변수를 가져왔습니다.", nil
}

func main() {
	lambda.Start(HandleRequest)
}
