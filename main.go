package main

import (
	"fmt"
	"groupie_tracker/autors"
	"html/template"
	"log"
	"net/http"
	"time"
	"strconv"
)

type PageData struct {
	TitleGroup	string
	Artists    	[]autors.Artist
	Long		[]int
}

var port = ":8080"

func main() {
	// Server routes
	http.HandleFunc("/", ArtHandler)
	/*http.HandleFunc("/artistinfo", ArtGetInfo)*/
	
	// Static file server for images and css files
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func ArtHandler(w http.ResponseWriter, req *http.Request) {
	autors.Find_Locations()
	var err error
	// Retrieving artist data about and details like dates: locations
	artists, err := autors.GetConcertDetails()
	if err != nil {
        log.Println("Erreur lors de la récupération des artistes:", err)
        http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
        return
    }

	// Retrieve the start and end dates chosen by the user
    startDateStr := req.URL.Query().Get("research-startDate")
    endDateStr := req.URL.Query().Get("research-endDate")

    if startDateStr != "" && endDateStr != "" {
        startDate, err := time.Parse("2006-01-02", startDateStr)
        if err != nil {
            http.Error(w, "Date de début invalide", http.StatusBadRequest)
            return
        }

        endDate, err := time.Parse("2006-01-02", endDateStr)
        if err != nil {
            http.Error(w, "Date de fin invalide", http.StatusBadRequest)
            return
        }

        // Filter artists by concert dates
        artists = autors.FilterArtistsByConcertDateRange(artists, startDate, endDate)
    }

	// Récupérer la sélection du nombre d'artistes
	numArtistsStr := req.URL.Query().Get("nombre")
	numArtists, err := strconv.Atoi(numArtistsStr)
	if err != nil || numArtists <= 0 {
		numArtists = len(artists) // Si pas de sélection valide, afficher tous les artistes
	}
	
	// Limiter la liste des artistes selon la sélection
	if numArtists < len(artists) {
		artists = artists[:numArtists]
	}

	var long []int
	for i := 1; i <= len(artists); i++ {
		long = append(long, i)
	}
	
	// Retrieving selected creation dates
    before1980 := req.URL.Query().Get("before-1980") != ""
    date1980to1990 := req.URL.Query().Get("1980-1990") != ""
    date1990to2000 := req.URL.Query().Get("1990-2000") != ""
    date2000to2010 := req.URL.Query().Get("2000-2010") != ""
    after2010 := req.URL.Query().Get("after 2010") != ""

    // Applying the filters
    if (before1980 || date1980to1990 || date1990to2000 || date2000to2010 || after2010) {
        artists = autors.FilterArtistsByCreationDates(artists, before1980, date1980to1990, date1990to2000, date2000to2010, after2010)
    }

	// Retrieve the search term
	searchTerm := req.URL.Query().Get("search")

	// If a search term is provided, filter artists, places and dates
	if searchTerm != "" {
		artists = autors.FilterArtistsBySearch(artists, searchTerm)
	}

	// Checking sorting parameters
	categorie := req.URL.Query().Get("categorie")
	if categorie == "reverseSens" {
		artists = autors.Filter_By_Name_Reversed(artists)
	} else {
		artists = autors.Filter_By_Name(artists)
	}

	categorie2 := req.URL.Query().Get("categorie2")
    if categorie2 == "reverseCreation" {
        artists = autors.Filter_By_Creation_Reversed(artists)
    } else if categorie2 == "normalCreation" {
        artists = autors.Filter_By_Creation(artists)
    }
	
	// Data for the template
	pageData := PageData{
		TitleGroup: "Groupie Trackers",
		Artists:    artists,
		Long:		long,
	}
	
	// Load and run the HTML template
	tmpl, err := template.New("home").ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors du chargement du template: %v", err), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du template: %v", err), http.StatusInternalServerError)
		return
	}
}

func ArtGetInfo(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		fmt.Println(req.Method)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	target := req.FormValue("artist")
	fmt.Println("artiste = ", target)
	artists, _ := autors.GetArtists()

	for _, artist := range artists {
		if artist.Name == target {
			tmpl, err := template.New("artistinfo").ParseFiles("templates/artistinfo.html")
			if err != nil {
				http.Error(w, fmt.Sprintf("Erreur lors du chargement du template: %v", err), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, artist)
			if err != nil {
				http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du template: %v", err), http.StatusInternalServerError)
			}
			return
		}
	}
	fmt.Println("Erreur lors du chargement de l'ariste, retour à la page principale")
	ArtHandler(w, req)	
}
