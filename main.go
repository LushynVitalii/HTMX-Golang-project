package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Привіт мій друже!")

	h1 := func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("index.html"))

		films :=map[string][]Film{
			"Films": {
				{Title: "First title", Director: "1st Director"},
				{Title: "Second title", Director: "2nd Director"},
				{Title: "Third title", Director: "3rd Director"},
			},
		}
		tmpl.Execute(w, films)

		// for learning purpose here
		// tmpl := template.Must(template.ParseFiles("index.html"))
		// tmpl.Execute(w, nil)

		// io.WriteString(w, "Hello YO")
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
