package autors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

func Length(url string) int {
	nb := 0
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	str := string(body)
	for i := 0; i <= len(str)-4; i++ {
		if str[i:i+4] == `"id"` {
			nb++
		}
	}
	return nb
}

func ReadLocation(body []byte) []location {
	var api map[string][]location

	err2 := json.Unmarshal([]byte(body), &api)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	index := api["index"]
	var art0 location //Création d'un artiste vide pour garder les Id en place
	index = append([]location{art0}, index...)
	return index
}

func OpenAllLocations() []location {
	url := "https://groupietrackers.herokuapp.com/api/locations"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return ReadLocation(body)
}
