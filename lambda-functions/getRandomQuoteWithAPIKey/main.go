package main

import (
	"log"

	"github.com/Skjaldbaka17/WhoTheFuckSaidThat/utils"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

func Handler(request Request) (utils.ResponseBody, error) {
	respBody, err := utils.GetRandomQuote()
	log.Println(respBody)
	if err != nil {
		return utils.ResponseBody{}, err
	}
	return respBody, nil
}

func main() {
	lambda.Start(Handler)
}
