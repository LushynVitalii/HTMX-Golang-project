package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type BetaUser struct {
	ImageUrl    string
	CompanyName string
	CompanyType string
	CompanyUrl  string
}

type PageContent struct {
	Header           string
	Tel              string
	H1               string
	Paragraph        string
	ProjectsTitle    string
	ProjectsSubTitle string
	BetaUsers        []BetaUser
}

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))
var err error

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}
	fmt.Println("Привіт мій друже! Port:", port, os.Getenv("FLY_REGION"))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/en" {
			data := PageContent{
				Header:           "For a Demo, Simply text « DEMO » at",
				Tel:              "1 (855) 950-4006",
				H1:               "Convert your missed calls into business opportunities",
				Paragraph:        "Coming soon to a phone system near you!",
				ProjectsTitle:    "Every. Call. Matters.",
				ProjectsSubTitle: "Create seamless engagement with your callers at all times. Whether you are open or closed, free to pick up the phone or busy, make every call count.",
				BetaUsers: []BetaUser{
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/jatoba.png", CompanyName: "Jatoba", CompanyType: "Restaurant", CompanyUrl: "https://www.jatobamontreal.com/"},
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/lallier.png", CompanyName: "Groupe Lallier", CompanyType: "Car Dealership", CompanyUrl: "https://www.lallier.com/en"}, {ImageUrl: "https://www.talksoon.com/assets/img/partners/bardagi.png", CompanyName: "Bardagi", CompanyType: "Real Estate", CompanyUrl: "https://www.bardagi.com/"},
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/golf-versant.png", CompanyName: "Golf le Versant", CompanyType: "Golf", CompanyUrl: "https://golfleversant.com/"},
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/vitre_net.png", CompanyName: "Vitre Net", CompanyType: "Services", CompanyUrl: "https://vitres.net/en/"},
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/germain.png", CompanyName: "Germain", CompanyType: "Hospitality", CompanyUrl: "https://www.germainhotels.com/en"},
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/sosticket.png", CompanyName: "SOS Ticket", CompanyType: "Legal", CompanyUrl: "https://www.sosticket.ca/"},
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/eco-odyssee.png", CompanyName: "Eco-Odyssee", CompanyType: "Recreation", CompanyUrl: "https://www.eco-odyssee.ca/en/"},
				},
			}

			err = t.ExecuteTemplate(w, "index_en.html", data)

		} else if r.URL.Path == "/fr" {
			data := PageContent{
				Header:           "Pour une démo, envoyez simplement le mot « DEMO » au",
				Tel:              "1 (855) 950-4006",
				H1:               "Transformez vos appels manqués en opportunités commerciales",
				Paragraph:        "Bientôt disponible sur un système téléphonique près de chez vous !",
				ProjectsTitle:    "Chaque appel compte.",
				ProjectsSubTitle: "Créez une interaction transparente avec vos appelants à tout moment. Que vous soyez ouverts ou fermés, disponibles pour prendre l'appel ou occupés, faites en sorte que chaque appel compte.",
				BetaUsers: []BetaUser{
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/jatoba.png", CompanyName: "Jatoba", CompanyType: "Restaurant", CompanyUrl: "https://www.jatobamontreal.com/"},
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/lallier.png", CompanyName: "Groupe Lallier", CompanyType: "Automobile", CompanyUrl: "https://www.lallier.com/en"},
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/bardagi.png", CompanyName: "Bardagi", CompanyType: "Immobilier", CompanyUrl: "https://www.bardagi.com/"},
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/golf-versant.png", CompanyName: "Golf le Versant", CompanyType: "Golf", CompanyUrl: "https://golfleversant.com/"},
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/vitre_net.png", CompanyName: "Vitre Net", CompanyType: "Services", CompanyUrl: "https://vitres.net/en/"},
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/germain.png", CompanyName: "Germain", CompanyType: "Hôtellerie", CompanyUrl: "https://www.germainhotels.com/en"},
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/sosticket.png", CompanyName: "SOS Ticket", CompanyType: "Juridique", CompanyUrl: "https://www.sosticket.ca/"},
					{ImageUrl: "https://www.talksoon.com/assets/img/partners/eco-odyssee.png", CompanyName: "Eco-Odyssee", CompanyType: "Loisirs", CompanyUrl: "https://www.eco-odyssee.ca/en/"},
				},
			}

			err = t.ExecuteTemplate(w, "index_fr.html", data)

		} else if r.URL.Path == "/toggle-menu" {
			err = t.ExecuteTemplate(w, "menu_modal.html", nil)
		} else {
			http.NotFound(w, r)
		}
		if err != nil {
			http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		}

	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
