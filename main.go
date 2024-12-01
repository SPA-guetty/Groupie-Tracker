package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var port = ":8080"

func main() {

	//Routes of server
	http.HandleFunc("/date", dateHandler)

	// Démarrer le serveur HTTP sur le port 8080
	fmt.Println("Server started on http://localhost:8080/date")
	log.Fatal(http.ListenAndServe(port, nil))
}

// Représentation d'une date de concert
type date struct {
	Id    string   `json:"id"`
	Dates []string `json:"date"`
}

// Fonction Handler concernant la fonctionnalité "date"
func dateHandler(w http.ResponseWriter, req *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/dates"

	// Effectuer la requête HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la requête APIDate: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Vérifier si la requête a réussi (code HTTP 200)
	if resp.StatusCode != 200 {
		http.Error(w, fmt.Sprintf("Erreur API, code statut: %d", resp.StatusCode), resp.StatusCode)
		return
	}

	// Lire le corps de la réponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la lecture de la réponse: %v", err), http.StatusInternalServerError)
		return
	}

	// Analyser les données JSON dans une variable de type date
	var data []date
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'analyse du JSON: %v", err), http.StatusInternalServerError)
		return
	}

	// Définir l'en-tête de la réponse HTTP comme étant du JSON
	w.Header().Set("Content-Type", "application/json")

	// Convertir la structure Go en JSON et l'envoyer en réponse
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'envoi de la réponse: %v", err), http.StatusInternalServerError)
	}
}
