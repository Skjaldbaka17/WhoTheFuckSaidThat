package utils

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
