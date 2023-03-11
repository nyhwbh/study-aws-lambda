package main

import (
	"context"
	"studylambda/savedata"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (string, error) {
	savedata.SaveDataAll()

	return "Google Spread Sheet를 업데이트 했습니다.", nil
}

func main() {
	lambda.Start(HandleRequest)
}
