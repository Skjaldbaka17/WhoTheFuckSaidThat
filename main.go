package main

import (
	"html/template"
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

	var templates = template.Must(template.ParseFiles("./static/index.html"))
	err = templates.ExecuteTemplate(rw, "index.html", respBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	s := http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets/")))
	r.PathPrefix("/assets/").Handler(s)

	err := http.ListenAndServe(":"+utils.GetEnvVariable("PORT"), r)
	if err != nil {
		panic(err)
	}
}
