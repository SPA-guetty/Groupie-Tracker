package locations

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

func ReadLocation(body []byte) {
	var artiststab location
	err2 := json.Unmarshal([]byte(body), &artiststab)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	fmt.Println(artiststab)
}

func OpenLocation(id string) {
	urlint, _ := strconv.Atoi(id)
	if (urlint < 1 || urlint > Length("https://groupietrackers.herokuapp.com/api/locations")) {
		fmt.Println("Error: locations index is out of range")
		return
	}
	url := "https://groupietrackers.herokuapp.com/api/locations/" + id
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	ReadLocation(body)
}