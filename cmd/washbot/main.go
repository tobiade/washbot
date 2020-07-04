package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tobiade/washbot"
)

func main() {
	lambda.Start(washbot.Handler)
}
