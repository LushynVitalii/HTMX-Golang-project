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
					{ImageUrl: "https://images.squarespace-cdn.com/content/v1/54614c38e4b0a75b67f464fc/1464793513567-008V7VS9I11GC3AAO3PU/JATOBA-201411-Logo-Dark-FINAL-Nov17.jpg?format=750w", CompanyName: "Jatoba", CompanyType: "Restaurant", CompanyUrl: "https://www.jatobamontreal.com/"},
					{ImageUrl: "https://img.sm360.ca/images/web/lallier/4216/lallier-logo-0-0-0-0-16680077541694393867928.png", CompanyName: "Groupe Lallier", CompanyType: "Car Dealership", CompanyUrl: "https://www.lallier.com/en"}, 
					{ImageUrl: "https://www.bardagi.com/app/uploads/2024/07/bardagi-fb-img.jpg", CompanyName: "Bardagi", CompanyType: "Real Estate", CompanyUrl: "https://www.bardagi.com/"},
					{ImageUrl: "https://golfleversant.com/wp-content/uploads/2024/07/cropped-logo_golf_le_versant_2024.png", CompanyName: "Golf le Versant", CompanyType: "Golf", CompanyUrl: "https://golfleversant.com/"},
					{ImageUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ3Jdk-4u4Y_7IeVR6uam1D3tct9qqVT5QyhA&s", CompanyName: "Vitre Net", CompanyType: "Services", CompanyUrl: "https://vitres.net/en/"},
					{ImageUrl: "https://assets.germainhotels.com/germain-website-assets-production/website-assets/Germain-H%C3%B4tels/Logo_GermainHotels_FR_2019_Black_Square.png", CompanyName: "Germain", CompanyType: "Hospitality", CompanyUrl: "https://www.germainhotels.com/en"},
					{ImageUrl: "https://images.prismic.io//intuzwebsite/59af8c5c-927b-4196-8128-be14a8af517f_osticket.png?w=1200&q=75&auto=format,compress&fm=png8", CompanyName: "SOS Ticket", CompanyType: "Legal", CompanyUrl: "https://www.sosticket.ca/"},
					{ImageUrl: "https://www.eco-odyssee.ca/wp-content/uploads/2022/09/a2b979_17b59a097c4c40f58b9f49cc5b785cb1_mv2.jpg.webp", CompanyName: "Eco-Odyssee", CompanyType: "Recreation", CompanyUrl: "https://www.eco-odyssee.ca/en/"},
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
					{ImageUrl: "https://images.squarespace-cdn.com/content/v1/54614c38e4b0a75b67f464fc/1464793513567-008V7VS9I11GC3AAO3PU/JATOBA-201411-Logo-Dark-FINAL-Nov17.jpg?format=750w", CompanyName: "Jatoba", CompanyType: "Restaurant", CompanyUrl: "https://www.jatobamontreal.com/"},
					{ImageUrl: "https://img.sm360.ca/images/web/lallier/4216/lallier-logo-0-0-0-0-16680077541694393867928.png", CompanyName: "Groupe Lallier", CompanyType: "Automobile", CompanyUrl: "https://www.lallier.com/en"},
					{ImageUrl: "https://www.bardagi.com/app/uploads/2024/07/bardagi-fb-img.jpg", CompanyName: "Bardagi", CompanyType: "Immobilier", CompanyUrl: "https://www.bardagi.com/"},
					{ImageUrl: "https://golfleversant.com/wp-content/uploads/2024/07/cropped-logo_golf_le_versant_2024.png", CompanyName: "Golf le Versant", CompanyType: "Golf", CompanyUrl: "https://golfleversant.com/"},
					{ImageUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ3Jdk-4u4Y_7IeVR6uam1D3tct9qqVT5QyhA&s", CompanyName: "Vitre Net", CompanyType: "Services", CompanyUrl: "https://vitres.net/en/"},
					{ImageUrl: "https://assets.germainhotels.com/germain-website-assets-production/website-assets/Germain-H%C3%B4tels/Logo_GermainHotels_FR_2019_Black_Square.png", CompanyName: "Germain", CompanyType: "Hôtellerie", CompanyUrl: "https://www.germainhotels.com/en"},
					{ImageUrl: "https://images.prismic.io//intuzwebsite/59af8c5c-927b-4196-8128-be14a8af517f_osticket.png?w=1200&q=75&auto=format,compress&fm=png8", CompanyName: "SOS Ticket", CompanyType: "Juridique", CompanyUrl: "https://www.sosticket.ca/"},
					{ImageUrl: "https://www.eco-odyssee.ca/wp-content/uploads/2022/09/a2b979_17b59a097c4c40f58b9f49cc5b785cb1_mv2.jpg.webp", CompanyName: "Eco-Odyssee", CompanyType: "Loisirs", CompanyUrl: "https://www.eco-odyssee.ca/en/"},
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
