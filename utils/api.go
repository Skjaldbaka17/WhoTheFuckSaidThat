package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Gets the random quote from the api
func GetRandomQuote() (ResponseBody, error) {
	url := GetEnvVariable("API_BASE_URL") + "/api/quotes/random"
	reqBody := PostRequestBody{ApiKey: GetEnvVariable("API_KEY")}
	var jsonStr = []byte(reqBody.ToString())
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if err != nil {
		log.Println("error setting request for api:", err)
		return ResponseBody{}, err
	}
	req.Header.Set("Content-Type", "application/json")

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
	err := godotenv.Load("./.env")
	if err != nil {
		log.Printf("Error loading .env file")
	}
	return os.Getenv(key)
}
