package main

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

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var tab []artist
	err := json.Unmarshal([]byte(body), &tab)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, e := range tab {
		fmt.Println(e.Name)
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