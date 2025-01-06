package autors

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"groupie_tracker/concertdates"
	"strconv"
)

type Concert struct {
    Date    	 string   `json:"date"`
    Location 	 string   `json:"location"`
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
	Concerts     []string `json:"concerts"`  //Stocke les dates et lieux de concerts
}

func GetArtists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	// Création de la requête HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Erreur lors de la création de la requête: %v", err)
		return nil, err
	}

	// Envoi de la requête et réception de la réponse
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Erreur lors de la réception de la réponse: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	// Vérifier si la réponse est un succès
	if res.StatusCode != http.StatusOK {
		log.Printf("Erreur: statut HTTP %d", res.StatusCode)
		return nil, fmt.Errorf("Erreur HTTP %d: %s", res.StatusCode, res.Status)
	}

	// Lire le corps de la réponse
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Erreur lors de la lecture du corps de la réponse: %v", err)
		return nil, err
	}

	// Décodage des données JSON dans une tranche d'artistes
	var artists []Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		log.Printf("Erreur lors du décodage JSON: %v", err)
		return nil, err
	}

	// Ensuite, vous allez associer les concerts à chaque artiste
	for i := range artists {
		// Récupérer les données des concerts concernant les artistes à partir de l'ID
		concertsData := concertdates.OpenDates(strconv.Itoa(artists[i].Id))
		var concerts []Concert
		for _, date := range concertsData.Dates {
			concerts = append(concerts)
			
			location := artists[i].Locations
			artists[i].Concerts = append(artists[i].Concerts, fmt.Sprintf("%s : %s", date, location))
		}
	}

	return artists, nil
}
