package handlerfolder

import (
	"fmt"
	"log"
	"net/http"
	"groupie_tracker/autors"
	"html/template"
	"time"
	"strconv"
)

type PageData struct {
	TitleGroup		string
	Artists    		[]autors.Artist
	Long			[]int
	Search			string
	Categorie		bool
	Categorie2		bool
	Croissant 		bool
	Before1980 		bool
    Date1980to1990 	bool
    Date1990to2000 	bool
    Date2000to2010 	bool
    After2010 		bool
	Actual			int
	Previous 		bool
	Next 			bool
}

func FindMethod(w http.ResponseWriter, req *http.Request, met string) string {
	if req.Method != "POST" {return ""}
	return req.FormValue(met)
}

func GetAnAmount(tab []autors.Artist, nbprint int, idprint	int) []autors.Artist {
	first := (nbprint*(idprint-1))
	last := nbprint*idprint
	if last > len(tab) {
		last = len(tab)
	}
	tab = tab[first:last]
	return tab
}

func NextAndPrevious(w http.ResponseWriter, req *http.Request) int {
	if req.Method != "POST" {return 1}
	switcher := req.FormValue("switch")
	actualvaluestr := req.FormValue("actualvalue")
	actualvalue, err := strconv.Atoi(actualvaluestr)
	if err != nil {
		fmt.Println("Error:", err)
		return 1
	}
	if switcher == "<" {
		actualvalue -= 1
	} else if switcher == ">" {
		actualvalue++
	}
	return actualvalue
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

	var actual int
	if FindMethod(w, req, "filters") == "Obtenir" {
		actual = 1
	} else {
		actual = NextAndPrevious(w,req)
	}

	var long []int
	for i := 1; i <= len(artists); i++ {
		long = append(long, i)
	}

	// Retrieve the start and end dates chosen by the user
    startDateStr := FindMethod(w, req, "research-startDate")
    endDateStr := FindMethod(w, req, "research-endDate")

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
	
	// Retrieving selected creation dates
    before1980 := FindMethod(w, req, "before-1980") != ""
    date1980to1990 := FindMethod(w, req, "1980-1990") != ""
    date1990to2000 := FindMethod(w, req, "1990-2000") != ""
    date2000to2010 := FindMethod(w, req, "2000-2010") != ""
    after2010 := FindMethod(w, req, "after 2010") != ""

    // Applying the filters
    if (before1980 || date1980to1990 || date1990to2000 || date2000to2010 || after2010) {
        artists = autors.FilterArtistsByCreationDates(artists, before1980, date1980to1990, date1990to2000, date2000to2010, after2010)
    }

	// Retrieve the search term
	searchTerm := FindMethod(w, req, "search")
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

	// Retrieve the selection of number of artists
	numArtistsStr := FindMethod(w, req, "nombre")
	//Show only the asked number of artists
	numArtists, err := strconv.Atoi(numArtistsStr)
	if err != nil || numArtists <= 0 || numArtists == 52 {
		numArtists = len(artists) // Si pas de sélection valide, afficher tous les artistes
		long = long[:51]
		longbis := []int{52}
		long = append(longbis, long...)
	} else if numArtists < len(long) {
		if numArtists < len(artists) {
			artists = GetAnAmount(artists, numArtists, actual)
		}
		long = append(long[:numArtists-1], long[numArtists:]...)
		longbis := []int{numArtists}
		long = append(longbis, long...)
	}

	// Data for the template
	pageData := PageData{
		TitleGroup: 	"Groupie Trackers",
		Artists:   		artists,
		Long:			long,
		Search:			searchTerm,
		Categorie:		(categorie != "reverseSens"),
		Categorie2:		(categorie2 != "reverseCreation"),
		Croissant: 		(categorie2 == "forgetCreation"),
		Before1980:		before1980,
		Date1980to1990:	date1980to1990,
		Date1990to2000:	date1990to2000,
		Date2000to2010:	date2000to2010,
		After2010:		after2010,
		Actual:			actual,
		Previous: 		(actual > 1),
		Next:			(len(long) > (numArtists*actual)),
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
