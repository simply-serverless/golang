package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(timeDisplayHandler)
}

func timeDisplayHandler() {
	fmt.Println(time.Now().Format(time.RFC850))
}
