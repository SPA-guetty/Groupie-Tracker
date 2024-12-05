package autors

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type artist struct {
	Id int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
	Locations string `json:"locations"`
	ConcertDates string `json:"concertDates"`
	Relations string `json:"relations"`
}

func GetArtists() []artist {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var tab []artist
	err := json.Unmarshal([]byte(body), &tab)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	for _, e := range tab {
		fmt.Println(e.Name)
	}
	return tab
}