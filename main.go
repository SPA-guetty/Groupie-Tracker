package main

import (
	"fmt"
	"groupietracker/autors"
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

	// Serveur de fichiers statiques pour les images et fichiers css
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func ArtHandler(w http.ResponseWriter, req *http.Request) {
	// Récupérer les données des artistes
	artists, err := autors.GetArtists()
	autors.Print_Locations()
	if err != nil {
		log.Fatalf("Erreur lors de la récupération des artistes: %v", err)
	}
	artists = autors.Filter_By_Name(artists)
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
