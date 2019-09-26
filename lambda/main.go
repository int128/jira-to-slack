package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	lambdaRouter "github.com/int128/jira-to-slack/pkg/router/lambda"
)

func main() {
	switch os.Getenv("LAMBDA_TYPE") {
	case "APIGateway":
		lambda.Start(lambdaRouter.APIGateway)
	case "ALBTargetGroup":
		lambda.Start(lambdaRouter.ALBTargetGroup)
	default:
		log.Fatal("you need to set LAMBDA_TYPE environment variable")
	}
}
