package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

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

func (body *PostRequestBody) toString() string {

	out, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func Home(rw http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("./tmpl/home.html"))
	url := "http://localhost:8080/api/quotes/random"
	fmt.Println("URL:>", url)

	reqBody := PostRequestBody{ApiKey: GetEnvVariable("API_KEY")}
	var jsonStr = []byte(reqBody.toString())
	log.Println(string(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var respBody ResponseBody
	json.NewDecoder(resp.Body).Decode(&respBody)
	log.Println(respBody)
	err = templates.ExecuteTemplate(rw, "home.html", respBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func GetEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load("./.env")

	if err != nil {
		log.Printf("Error loading .env file")
		err = godotenv.Load("../.env")
		if err != nil {
			log.Printf("Error loading ../.env file")
		}
	}

	return os.Getenv(key)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	s := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	r.PathPrefix("/assets/").Handler(s)

	err := http.ListenAndServe(":"+GetEnvVariable("PORT"), r)

	if err != nil {
		panic(err)
	}
}
