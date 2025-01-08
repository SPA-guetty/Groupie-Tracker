package autors

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"groupie_tracker/concertdates"
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
	body, _ := io.ReadAll(res.Body)
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
	var art0 location //Creating an empty artist to keep the Id in place
	index = append([]location{art0}, index...)
	return index
}

func OpenAllLocations() []location {
	url := "https://groupietrackers.herokuapp.com/api/locations"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return ReadLocation(body)
}

func GetConcertDatesAndLocations(artistId int) map[string]string {
	datesData := concertdates.OpenDates(strconv.Itoa(artistId))
	locationsData := OpenAllLocations()

	// Allows to find the location according to each concert date for each artist (artistId)
	dateLocationMap := make(map[string]string)
	for i, date := range datesData.Dates {
		// Association of concert date and venue for each artist
		if i < len(locationsData) && locationsData[i].Id == artistId {
			dateLocationMap[date] = locationsData[i].Locations[i]
		}
	}

	return dateLocationMap
}