package main

import (
	"encoding/json"
	"fmt"
	"groupitracker/dates"
	"io"
	"log"
	"net/http"
	"strconv"
)

var port = ":8080"

func main() {
	dates.OpenDates("52")
	//Routes of server
	http.HandleFunc("/date", dateHandler)

	// Démarrer le serveur HTTP sur le port 8080
	fmt.Println("Server started on http://localhost:8080/date")
	log.Fatal(http.ListenAndServe(port, nil))

}

// Représentation d'une date de concert
type date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

var datatab []date

// Fonction Handler concernant la fonctionnalité "date"
func dateHandler(w http.ResponseWriter, req *http.Request) {

	for i := 1; i <= 52; i++ {
		url := "https://groupietrackers.herokuapp.com/api/dates/" + strconv.Itoa(i)

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
		var data date

		//fmt.Println("test", string(body))
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors de l'analyse du JSON: %v", err), http.StatusInternalServerError)
			return
		}
		datatab = append(datatab, data)
	}

	// Définir l'en-tête de la réponse HTTP comme étant du JSON
	w.Header().Set("Content-Type", "application/json")

	// Convertir la structure Go en JSON et l'envoyer en réponse
	if err := json.NewEncoder(w).Encode(datatab); err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'envoi de la réponse: %v", err), http.StatusInternalServerError)
	}
}
