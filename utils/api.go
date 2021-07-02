package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Gets the random quote from the api
func GetRandomQuote() (ResponseBody, error) {
	url := GetEnvVariable("API_BASE_URL") + "/quotes/random"
	var jsonStr = []byte("{}")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if err != nil {
		log.Println("error setting request for api:", err)
		return ResponseBody{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", GetEnvVariable("API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error getting random quote from api:", err)
		return ResponseBody{}, err
	}
	defer resp.Body.Close()

	var respBody ResponseBody
	json.NewDecoder(resp.Body).Decode(&respBody)
	return respBody, nil
}

func GetEnvVariable(key string) string {
	return os.Getenv(key)
}
