package utils

import "encoding/json"

type PostRequestBody struct {
	ApiKey       string `json:"apiKey"`
	AuthorId     int    `json:"authorId"`
	Language     string `json:"language"`
	SearchString string `json:"searchString"`
	TopicId      int    `json:"topicId"`
}

type ResponseBody struct {
	AuthorId    int    `json:"authorId"`
	IsIcelandic bool   `json:"isIcelandic"`
	Name        string `json:"name"`
	Quote       string `json:"quote"`
	QuoteId     int    `json:"quoteId"`
	TopicId     int    `json:"topicId"`
	TopicName   string `json:"topicName"`
	Message     string `json:"message"`
}

func (body *PostRequestBody) ToString() string {
	out, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	return string(out)
}
