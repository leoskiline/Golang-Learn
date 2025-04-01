package main

import (
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Nome  string
	Email string
}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := usuario{
			Nome:  "Leonardo",
			Email: "leonardo@hotmail.com",
		}

		templates.ExecuteTemplate(w, "home.html", u)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
