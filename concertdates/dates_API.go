package concertdates

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Representation of a concert date
type date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
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

func ReadDates(body []byte) date {
	// Analysis of JSON data in a date variable
	var data date
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Erreur lors de l'analyse du JSON: %v", err)
	}
	for index, date := range data.Dates {
		cleandate := ""
		for _, run := range (date) {
			if string(run) != "*" {
				if string(run) == "-" {
					cleandate += " "
				} else {
					cleandate += string(run)
				}
			}
		}
		data.Dates[index] = cleandate
	}
	return data
}

func OpenDates(id string) date {
	urlint, _ := strconv.Atoi(id)
	if urlint < 1 || urlint > Length("https://groupietrackers.herokuapp.com/api/dates") {
		log.Fatal("Error: dates index is out of range")
	}
	url := "https://groupietrackers.herokuapp.com/api/dates/" + id
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Erreur lors de la requête HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful (HTTP code 200)
	if resp.StatusCode != 200 {
		log.Fatalf("La requête a échoué avec le statut: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du corps de la réponse: %v", err)
	}
	return ReadDates(body)
}

func Get_All_Dates() []date {
	var tab []date
	for i := 1; i <= Length("https://groupietrackers.herokuapp.com/api/dates"); i++ {
		val := strconv.Itoa(i)
		tab = append(tab, OpenDates(val))
	}
	return tab
}
