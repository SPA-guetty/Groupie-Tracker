package autors

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"groupie_tracker/concertdates"
	"strconv"
	"log"
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

func CleanLocations(tab []location) []location {
	for index, locations := range tab {
		var newtab []string
		for _, location := range locations.Locations {
			upper := true
			newlocation := ""
			for _, run := range location {
				if run == '_' {
					newlocation += " "
					upper = true
				} else if run == '-' {
					newlocation += " ("
					upper = true
				} else {
					if upper {
						newlocation += string(byte(run)-32)
						upper = false
					} else {
						newlocation += string(run)
					}
				}
			}
			newlocation += ")"
			newtab = append(newtab, newlocation)
		}
		tab[index].Locations = newtab
	}
	return tab
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
	index = CleanLocations(index)
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

	// Permet de trouver le lieu en fonction de chaque date de concert pour chaque artiste (artistId)
	dateLocationMap := make(map[string]string)
	for i, date := range datesData.Dates {
		// Association de la date de concert avec son lieu pour chaque artistId
		if i < len(locationsData) && locationsData[i].Id == artistId {
			dateLocationMap[date] = locationsData[i].Locations[i]
		}
	}
	fmt.Println(dateLocationMap)
	return dateLocationMap
}

func AssociateConcertsWithLocations() []Artist {
    // Récupérer les artistes, les lieux et les dates de concerts
    artists, err := GetArtists()
    if err != nil {
        log.Println("Erreur lors de la récupération des artistes:", err)
        return nil
    }

    locations := OpenAllLocations()
    concertDates := concertdates.Get_All_Dates()

    // Associer les dates aux lieux
    for i := range artists {
        var concertInfo []string
        for j := 0; j < len(concertDates); j++ {
            if len(concertDates[j].Dates) > j && len(locations) > j { // Vérifier les limites
                // Associer la date de concert au lieu de concert en fonction de leur indice
                concertInfo = append(concertInfo, fmt.Sprintf("%s : %s", concertDates[j].Dates[j], locations[j].Locations[j]))
            }
        }
        artists[i].ConcertDatesLocations = concertInfo // Sauvegarder les associations
    }
    return artists
}