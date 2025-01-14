package autors

import (
	"encoding/json"
	"fmt"
	"groupie_tracker/concertdates"
	"io"
	"log"
	"net/http"
	"time"
)

type Artist struct {
	Id                    int          `json:"id"`
	Image                 string       `json:"image"`
	Name                  string       `json:"name"`
	Members               []string     `json:"members"`
	CreationDate          int          `json:"creationDate"`
	FirstAlbum            string       `json:"firstAlbum"`
	Locations             string       `json:"locations"`
	Relations             string       `json:"relations"`
	ConcertDates          ConcertDates `json:"concertDates"`
	ConcertLocations      []string     `json:"concertLocations"`
	ConcertDatesLocations []string     `json:"concertDatesLocations"`
}

// Type customization for concert dates because there was a problem with dates that were not of []string type
type ConcertDates []string

// Function allowing customization of UnmarshalJSON as well to avoid any problems later between the GO code and the JSON decoding
func (cd *ConcertDates) UnmarshalJSON(data []byte) error {
	// Decoding as a table test
	var arrayDates []string
	if err := json.Unmarshal(data, &arrayDates); err == nil {
		*cd = arrayDates
		return nil
	}

	// If it is not an array, try decoding as a string
	var singleDate string
	if err := json.Unmarshal(data, &singleDate); err == nil {
		*cd = []string{singleDate}
		return nil
	}

	// If both fail, we return an error
	return fmt.Errorf("concertDates doit être soit un tableau de chaînes, soit une chaîne unique")
}

func GetArtists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	// Creating the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Erreur lors de la création de la requête: %v", err)
		return nil, err
	}

	// Sending the request and receiving the response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Erreur lors de la réception de la réponse: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	// Verification of success of response
	if res.StatusCode != http.StatusOK {
		log.Printf("Erreur: statut HTTP %d", res.StatusCode)
		return nil, fmt.Errorf("Erreur HTTP %d: %s", res.StatusCode, res.Status)
	}

	// Reading the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Erreur lors de la lecture du corps de la réponse: %v", err)
		return nil, err
	}

	// At this place we decode the JSON data in a slice of artists
	var artists []Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		log.Printf("Erreur lors du décodage JSON: %v", err)
		return nil, err
	}

	return artists, nil
}

func GetConcertDetails() ([]Artist, error) {
	// We pick up the artists
	artists, err := GetArtists()
	if err != nil {
		return nil, err
	}

	// We retrieve dates and locations
	dates := concertdates.Get_All_Dates()
	locations := OpenAllLocations()

	// Dates and places associated with each artist
	for i := range artists {
		// For each artist, we retrieve their concert dates and locations
		artistDates := dates[i].Dates
		artistLocations := []string{}
		for _, location := range locations {
			if location.Id == artists[i].Id {
				artistLocations = location.Locations
				break
			}
		}

		// Creation of a table for the storage of results in the form "date: location"
		var concertDetails []string
		for j := 0; j < len(artistDates) && j < len(artistLocations); j++ {
			concertDetails = append(concertDetails, artistDates[j] + " : " + artistLocations[j])
		}
		artists[i].ConcertDates = artistDates
		artists[i].ConcertLocations = concertDetails
	}

	return artists, nil
}

// Function to filter artists according to date range
func FilterArtistsByConcertDateRange(artists []Artist, startDate, endDate time.Time) []Artist {
    var filteredArtists []Artist

    // For each artist, check if at least one of their concert dates is in the meantime
    for _, artist := range artists {
        var hasConcertInRange bool
        for _, concertDateStr := range artist.ConcertDates {
            // Convertir the concert date in type : time.Time
            concertDate, err := time.Parse("02 01 2006", concertDateStr)
            if err != nil {
                log.Printf("Erreur lors de la conversion de la date du concert: %v", err)
                continue
            }

            if concertDate.After(startDate) && concertDate.Before(endDate) {
                hasConcertInRange = true
                break
            }
        }

        if hasConcertInRange {
            filteredArtists = append(filteredArtists, artist)
        }
    }

    return filteredArtists
}