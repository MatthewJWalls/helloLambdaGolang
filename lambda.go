package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

// json in
type Request struct {
	Message string `json:"message"`
}

// json out
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func Handler(request Request) (Response, error) {
	return Response{
		Message: "you sent the message " + request.Message,
		Ok: true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
