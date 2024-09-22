package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type BetaUser struct {
	ImageUrl    string
	CompanyName string
	CompanyType string
	CompanyUrl string
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

func main() {
	fmt.Println("Привіт мій друже!")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	homePage := func(w http.ResponseWriter, r *http.Request) {
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
			BetaUsers: []BetaUser{
				{ImageUrl: "https://www.talksoon.com/assets/img/partners/jatoba.png", CompanyName: "Jatoba", CompanyType: "Restaurant", CompanyUrl: "https://www.jatobamontreal.com/"},
				{ImageUrl: "https://www.talksoon.com/assets/img/partners/lallier.png", CompanyName: "Groupe Lallier", CompanyType: "Car Dealership", CompanyUrl: "https://www.lallier.com/en"},
				{ImageUrl: "https://www.talksoon.com/assets/img/partners/bardagi.png", CompanyName: "Bardagi", CompanyType: "Real Estate", CompanyUrl: "https://www.bardagi.com/"},
				{ImageUrl: "https://www.talksoon.com/assets/img/partners/golf-versant.png", CompanyName: "Golf le Versant", CompanyType: "Golf", CompanyUrl: "https://golfleversant.com/"},
				{ImageUrl: "https://www.talksoon.com/assets/img/partners/vitre_net.png", CompanyName: "Vitre Net", CompanyType: "Services", CompanyUrl: "https://vitres.net/en/"},
				{ImageUrl: "https://www.talksoon.com/assets/img/partners/germain.png", CompanyName: "Germain", CompanyType: "Hospitality", CompanyUrl: "https://www.germainhotels.com/en"},
				{ImageUrl: "https://www.talksoon.com/assets/img/partners/sosticket.png", CompanyName: "SOS Ticket", CompanyType: "Legal", CompanyUrl: "https://www.sosticket.ca/"},
				{ImageUrl: "https://www.talksoon.com/assets/img/partners/eco-odyssee.png", CompanyName: "Eco-Odyssee", CompanyType: "Recreation", CompanyUrl: "https://www.eco-odyssee.ca/en/"},
			},
		}
		tmpl.Execute(w, content)
	}
	http.HandleFunc("/en", homePage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
