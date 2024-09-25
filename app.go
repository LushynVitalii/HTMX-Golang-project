package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
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

func main() {
	port := os.Getenv("PORT")
    if port == "" {
        port = "8081"
    }

	fmt.Println("Привіт мій друже! Port:", port)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	homePage := func(w http.ResponseWriter, r *http.Request) {
		var tmpl *template.Template
		var content PageContent

		// Check if the user requests the French page
		switch r.URL.Path {
		case "/fr":
			tmpl, _ = template.ParseFiles("templates/index_fr.html")
			content = PageContent{
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
		case "/en", "/":
			tmpl, _ = template.ParseFiles("templates/index_en.html")
			content = PageContent{
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
		default:
			http.NotFound(w, r)
			return
		}

		tmpl.Execute(w, content)
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/en", homePage)
	http.HandleFunc("/fr", homePage)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
