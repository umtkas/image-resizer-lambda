package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/umtkas/image-resizer-lambda/internal/aws"
)

func main() {
	lambda.Start(aws.Handler)
}
