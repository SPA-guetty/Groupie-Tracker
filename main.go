package main

import (
	"fmt"
	"groupie_tracker/autors"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	TitleGroup string
	Artists    []autors.Artist
}

var port = ":8080"

func main() {
	// Routes du serveur
	http.HandleFunc("/", ArtHandler)
	http.HandleFunc("/artistinfo", ArtGetInfo)

	// Serveur de fichiers statiques pour les images et fichiers css
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func ArtHandler(w http.ResponseWriter, req *http.Request) {
	// Récupérer les données des artistes
	artists, err := autors.GetArtists()

	// Vérifier le paramètre de tri
	categorie := req.URL.Query().Get("categorie")
	if categorie == "reverseSens" {
		artists = autors.Filter_By_Name_Reversed(artists)
	} else {
		artists = autors.Filter_By_Name(artists)
	}

	categorie2 := req.URL.Query().Get("categorie2")
    if categorie2 == "reverseCreation" {
        artists = autors.Filter_By_Creation_Reversed(artists)
    } else {
        artists = autors.Filter_By_Creation(artists)
    }

	// Données pour le template
	pageData := PageData{
		TitleGroup: "Groupie Trackers",
		Artists:    artists,
	}

	// Charger et exécuter le template HTML
	tmpl, err := template.New("home").ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors du chargement du template: %v", err), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du template: %v", err), http.StatusInternalServerError)
	}
}

func ArtGetInfo(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		fmt.Println(req.Method)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	target := req.FormValue("artist")
	fmt.Println("ariste = ", target)
	artists, _ := autors.GetArtists()

	for _, artist := range artists {
		if artist.Name == target {
			tmpl, err := template.New("home").ParseFiles("templates/home.html")
			if err != nil {
				http.Error(w, fmt.Sprintf("Erreur lors du chargement du template: %v", err), http.StatusInternalServerError)
				return
			}
			fmt.Println(artist)
			err = tmpl.Execute(w, artist)
			if err != nil {
				http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du template: %v", err), http.StatusInternalServerError)
			}
			return
		}
	}
	fmt.Println("Erreur lors du chargement de l'ariste, retour à la page principale")
	ArtHandler(w, req)	
}
