package dates

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type date struct {
	Id    string   `json:"id"`
	Dates []string `json:"date"`
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

func ReadDates(body []byte) {
	var artiststab date
	err2 := json.Unmarshal([]byte(body), &artiststab)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	fmt.Println(artiststab)
}

func OpenDates(id string) {
	urlint, _ := strconv.Atoi(id)
	if urlint < 1 || urlint > Length("https://groupietrackers.herokuapp.com/api/locations") {
		fmt.Println("Error: locations index is out of range")
		return
	}
	url := "https://groupietrackers.herokuapp.com/api/locations/" + id
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	ReadDates(body)
}

// GetDatesByArtistID récupère toutes les dates de concerts d'un artiste par son ID
func GetDatesByArtist(Id string) date {
	var result date
	for _, date := range dates {
		if date.Id == Id {
			result = append(result, date)
		}
	}
	return result
}

// GetDates récupère toutes les dates de concert associées à un artiste via son ArtistID
func GetDates(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de l'artiste à partir de l'URL
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Artist ID is required", http.StatusBadRequest)
		return
	}

	// Récupérer les dates de concert associées à l'artiste
	dates := GetDatesByArtist(Id)

	// Retourner les dates de concert en format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dates)
}

