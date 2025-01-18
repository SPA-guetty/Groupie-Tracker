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
	Categorie	bool
	Categorie2	bool
	Croissant 	bool
	StartDate	string
	EndDate		string
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

func FindMethod(w http.ResponseWriter, req *http.Request, met string) string {
	if req.Method == "POST" {
		Answer := req.FormValue(met)
		fmt.Println("categorie:",Answer)
		return Answer
	}
	return ""
}

func ArtHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	autors.Find_Locations()
	var err error
	// Retrieving artist data about and details like dates: locations
	artists, err := autors.GetConcertDetails()
	if err != nil {
        log.Println("Erreur lors de la récupération des artistes:", err)
        http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
        return
    }
	
	var long []int
	for i := 1; i <= len(artists); i++ {
		long = append(long, i)
	}

	artists = autors.Filter_By_Name(artists)

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

	// Retrieve the selection of number of artists
	numArtistsStr := FindMethod(w, req, "nombre")
	
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
	categorie := FindMethod(w, req, "categorie")
	if categorie == "reverseSens" {
		artists = autors.Filter_By_Name_Reversed(artists)
	} else {
		artists = autors.Filter_By_Name(artists)
	}

	categorie2 := FindMethod(w, req, "categorie2")
    if categorie2 == "reverseCreation" {
        artists = autors.Filter_By_Creation_Reversed(artists)
    } else if categorie2 == "normalCreation" {
        artists = autors.Filter_By_Creation(artists)
    }
	
	//Show only the asked number of artists
	numArtists, err := strconv.Atoi(numArtistsStr)
	if err != nil || numArtists <= 0 || numArtists == 52 {
		numArtists = len(artists) // Si pas de sélection valide, afficher tous les artistes
		long = long[:51]
		longbis := []int{52}
		long = append(longbis, long...)
	}
	
	if numArtists < len(artists) {
		artists = artists[:numArtists]
		long = append(long[:numArtists-1], long[numArtists:]...)
		longbis := []int{numArtists}
		long = append(longbis, long...)
	}

	// Data for the template
	pageData := PageData{
		TitleGroup: "Groupie Trackers",
		Artists:    artists,
		Long:		long,
		Categorie:	(categorie != "reverseSens"),
		Categorie2:	(categorie2 != "reverseCreation"),
		Croissant: 	(categorie2 == "forgetCreation"),
	}
	fmt.Println(pageData.Categorie, categorie)
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
