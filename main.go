package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Skjaldbaka17/WhoTheFuckSaidThat/utils"
	"github.com/gorilla/mux"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	respBody, err := utils.GetRandomQuote()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var templates = template.Must(template.ParseFiles("./tmpl/home.html"))
	err = templates.ExecuteTemplate(rw, "home.html", respBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	s := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
	r.PathPrefix("/assets/").Handler(s)

	err := http.ListenAndServe(":"+utils.GetEnvVariable("PORT"), r)
	log.Println("HERE", utils.GetEnvVariable("PORT"))
	if err != nil {
		panic(err)
	}
}
