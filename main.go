package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

func length(url string) int {
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

func main() {
	argument := os.Args
	url := argument[1]
	urlint, _ := strconv.Atoi(url)
	if (urlint < 1 || urlint > length("https://groupietrackers.herokuapp.com/api/locations")) {
		fmt.Println("Error: locations index is out of range")
		return
	}
	url = "https://groupietrackers.herokuapp.com/api/locations/" + url
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var artiststab location
	err2 := json.Unmarshal([]byte(body), &artiststab)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	fmt.Println(artiststab)
}
