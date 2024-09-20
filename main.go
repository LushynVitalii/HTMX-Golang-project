package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type PageContent struct {
	Header           string
	Tel              string
	H1               string
	Paragraph        string
	ProjectsTitle    string
	ProjectsSubTitle string
}

func main() {
	fmt.Println("Привіт мій друже!")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatal("Error parsing template: ", err)
		}

		content := PageContent{
			Header:           "For a Demo, Simply text « DEMO » at",
			Tel:              "1 (855) 950-4006",
			H1:               "Convert your missed calls into business opportunities",
			Paragraph:        "Coming soon to a phone system near you!",
			ProjectsTitle:    "Every. Call. Matters.",
			ProjectsSubTitle: "Create seamless engagement with your callers at all times. Whether you are open or closed, free to pick up the phone or busy, make every call count.",
		}
		tmpl.Execute(w, content)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
