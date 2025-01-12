package concertdates

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"fmt"
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
	fmt.Println(data, "COUCOU")
	return data
}

func OpenDates(id string) date {
	fmt.Println("hi")
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
/*
func Get_All_Dates2() []date {
	var tab []date
	for i := 1; i <= Length("https://groupietrackers.herokuapp.com/api/dates"); i++ {
		val := strconv.Itoa(i)
		tab = append(tab, OpenDates2(val))
	}
	return tab
}*/

func Clean_Date(api []date) []date {
	for indexapi, dateindex := range api {
		for indexdate, date := range dateindex.Dates {
			cleandate := ""
			for _, run := range date {
				if string(run) != "*" {
					if string(run) == "-" {
						cleandate += " "
					} else {
						cleandate += string(run)
					}
				}
			}
			api[indexapi].Dates[indexdate] = cleandate
		}
	}
	return api
}

func Open_All_Dates(body []byte) []date {
	var api map[string][]date
	err := json.Unmarshal([]byte(body), &api)
	if err != nil {
		fmt.Println("Error:", err)
	}
	index := api["index"]
	index = Clean_Date(index)
	var date0 date //Date vide pour régler l'API et garder les id
	index = append([]date{date0}, index...)
	return index
}

func Get_All_Dates() []date {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return Open_All_Dates(body)
}
