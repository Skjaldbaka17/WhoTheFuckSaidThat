package main

import (
	"github.com/Skjaldbaka17/WhoTheFuckSaidThat/utils"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

type Response struct {
	StatusCode int                `json:"statusCode"`
	Headers    map[string]string  `json:"headers"`
	Body       utils.ResponseBody `json:"body"`
}

func Handler(request Request) (utils.ResponseBody, error) {
	respBody, err := utils.GetRandomQuote()
	if err != nil {
		return utils.ResponseBody{}, err
	}
	// headers := map[string]string{
	// 	// "Access-Control-Allow-Headers": "Content-Type",
	// 	"Access-Control-Allow-Origin": "*",
	// 	// "Access-Control-Allow-Methods": "OPTIONS,POST,GET",
	// }
	// return Response{
	// 	StatusCode: 200,
	// 	Headers:    headers,
	// 	Body:       respBody,
	// }, nil
	return respBody, nil
}

func main() {
	lambda.Start(Handler)
}
