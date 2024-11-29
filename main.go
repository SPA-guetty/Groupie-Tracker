package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

func main() {
	argument := os.Args
	url := argument[1]
	url = "https://groupietrackers.herokuapp.com/api/locations/" + url
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var artiststab location
	err := json.Unmarshal([]byte(body), &artiststab)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(artiststab)
}
