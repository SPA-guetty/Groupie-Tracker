package dates

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Représentation d'une date de concert
type date struct {
	Id    string   `json:"id"`
	Dates []string `json:"date"`
}

func Length(url string) int {
	nb := 0
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	str := string(body)
	for i := 0; i <= len(str)-4; i++ {
		if str[i:i+4] == `"id"` {
			nb++
		}
	}
	return nb
}

func ReadDates(body []byte) {
	var datatab date
	err := json.Unmarshal([]byte(body), &datatab)
	if err != nil {
		log.Fatalf("Erreur lors de l'analyse du JSON: %v", err)
	}
	// Afficher les résultats
	var data = []date{}
	for _, item := range data {
		fmt.Printf("Id: %d\n", item.Id)
		fmt.Println("Dates:")
		for _, date := range item.Dates {
			fmt.Println(date)
		}
	}
}

func OpenDates(id string) {
	urlint, _ := strconv.Atoi(id)
	if urlint < 1 || urlint > Length("https://groupietrackers.herokuapp.com/api/locations") {
		fmt.Println("Error: dates index is out of range")
		return
	}
	url := "https://groupietrackers.herokuapp.com/api/locations/" + id
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erreur lors de la requête HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Vérifier si la requête a réussi (code HTTP 200)
	if resp.StatusCode != 200 {
		log.Fatalf("La requête a échoué avec le statut: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du corps de la réponse: %v", err)
	}
	ReadDates(body)
}
