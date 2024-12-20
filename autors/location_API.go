package autors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

func ReadLocation(body []byte) location {
	var artist location
	err2 := json.Unmarshal([]byte(body), &artist)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	return artist
}

func OpenLocation(idint int) location {
	id := strconv.Itoa(idint)
	if (idint < 1 || idint > Length("https://groupietrackers.herokuapp.com/api/locations")) {
		fmt.Println("Error: locations index is out of range")
	}
	url := "https://groupietrackers.herokuapp.com/api/locations/" + id
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return ReadLocation(body)
}