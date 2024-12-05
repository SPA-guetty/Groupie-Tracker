package main

import (
	"encoding/json"
	"groupie_tracker/autors"
	"log"
	"net/http"
	"fmt"
)

func main() {
	log.Println("Listening on :8080...")
	http.HandleFunc("/", ArtHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ArtHandler(w http.ResponseWriter, req *http.Request) {
	tab := autors.GetArtists()

	fmt.Println("tab:",tab)
	
	w.Header().Set("Content-Type", "application/json")

	// Convertir la structure Go en JSON et l'envoyer en réponse
	if err := json.NewEncoder(w).Encode(tab); err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'envoi de la réponse: %v", err), http.StatusInternalServerError)
	}
}

/*
{"id":52,
"image":"https://groupietrackers.herokuapp.com/api/images/thechainsmokers.jpeg",
"name":"The Chainsmokers",
"members":["Alexander Pall","Andrew Taggart","Matt McGuire","Tony Ann"],
"creationDate":2008,
"firstAlbum":"15-03-2014",
"locations":"https://groupietrackers.herokuapp.com/api/locations/52",
"concertDates":"https://groupietrackers.herokuapp.com/api/dates/52",
"relations":"https://groupietrackers.herokuapp.com/api/relation/52"}
*/