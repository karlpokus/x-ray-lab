package main

import (
	"context"
	"io/ioutil"
	"time"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

var client = xray.Client(&http.Client{
	Timeout: 5 * time.Second, // tcp ttl
})

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(parent context.Context) (Response, error) {
	req, _ := http.NewRequest("GET", "https://ip.mrfriday.com", nil)
	ctx, cancel := context.WithTimeout(parent, 3*time.Second) // http ttl
	defer cancel()
	res, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return Response{
			StatusCode: 500,
			Body: err.Error(),
		}, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{
			StatusCode: 500,
			Body: err.Error(),
		}, err
	}
	return Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(b),
		Headers: map[string]string{
			"Content-Type":           "text/plain",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
