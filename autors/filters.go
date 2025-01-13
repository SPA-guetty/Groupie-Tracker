package autors

import (
	"groupie_tracker/concertdates"
	"strconv"
	"strings"
)

func Filter_By_Name(tab []Artist) []Artist {
	var new_tab []Artist
	for range tab {
		firstval := tab[0]
		firstint := 0
		for i, e := range tab {
			if e.Name < firstval.Name {
				firstval = e
				firstint = i
			}
		}
		new_tab = append(new_tab, firstval)
		tab = append(tab[:firstint], tab[(firstint+1):]...)
	}
	return new_tab
}

func Filter_By_Name_Reversed(tab []Artist) []Artist {
	var new_tab []Artist
	for range tab {
		firstval := tab[0]
		firstint := 0
		for i, e := range tab {
			if e.Name > firstval.Name {
				firstval = e
				firstint = i
			}
		}
		new_tab = append(new_tab, firstval)
		tab = append(tab[:firstint], tab[(firstint+1):]...)
	}
	return new_tab
}

func Filter_By_Creation(tab []Artist) []Artist {
	var new_tab []Artist
	for len(tab) > 0 {
		var tabart []Artist
		var tabint []int
		tabart = append(tabart, tab[0])
		tabint = append(tabint, 0)
		for i, e := range tab {
			if i > 0 {
				if e.CreationDate < tabart[0].CreationDate {
					tabart = []Artist{e}
					tabint = []int{i}
				} else if e.CreationDate == tabart[0].CreationDate {
					tabart = append(tabart, e)
					tabint = append(tabint, i)
				}
			}
		}
		
		for i := range tabart {
			new_tab = append(new_tab, tabart[i])
			tab = append(tab[:(tabint[i]-i)], tab[(tabint[i]-i+1):]...)
		}
	}
	return new_tab
}

// Sorts artists by creation date in descending order
func Filter_By_Creation_Reversed(tab []Artist) []Artist {
    var new_tab []Artist
    for len(tab) > 0 {
        var tabart []Artist
        var tabint []int
        tabart = append(tabart, tab[0])
        tabint = append(tabint, 0)
        for i, e := range tab {
            if i > 0 {
                if e.CreationDate > tabart[0].CreationDate {
                    tabart = []Artist{e}
                    tabint = []int{i}
                } else if e.CreationDate == tabart[0].CreationDate {
                    tabart = append(tabart, e)
                    tabint = append(tabint, i)
                }
            }
        }
        for i := range tabart {
            new_tab = append(new_tab, tabart[i])
            tab = append(tab[:(tabint[i]-i)], tab[(tabint[i]-i+1):]...)
        }
    }
    return new_tab
}

func Find_date(target string, tab []Artist) []Artist {
	var newtab []Artist
	for _, artist := range tab {
		data := concertdates.OpenDates(strconv.Itoa(artist.Id))
		for _, date := range data.Dates {
			if date == target {
				newtab = append(newtab, artist)
			}
		}
	}
	return newtab
}

func Invert_Dates(date string) string {
	var tab []string
	tab = append(tab, "")
	tabint := 0
	for i := 0; i < len(date); i++ {
		if date[i] == '-' {
			tabint = tabint + 1
			tab = append(tab, "")
		} else if date[i] != '*' {
			tab[tabint] = tab[tabint] + string(date[i])
		}
	}
	return string(tab[2] + tab[1] + tab[0])
}

func Get_Between_Dates(targetmin string, targetmax string, tab []Artist) []Artist {
	targetmin = Invert_Dates(targetmin)
	targetmax = Invert_Dates(targetmax)
	var newtab []Artist
	for _, artist := range tab {
		data := concertdates.OpenDates(strconv.Itoa(artist.Id))
		for _, date := range data.Dates {
			date = Invert_Dates(date)
			if date >= targetmin && date <= targetmax {
				newtab = append(newtab, artist)
				break
			}
		}
	}
	return newtab
}

func FilterArtistsBySearch(artists []Artist, searchTerm string) []Artist {
	var filteredArtists []Artist
	searchTerm = strings.ToLower(searchTerm)

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), searchTerm) {
			filteredArtists = append(filteredArtists, artist)
			continue
		}

		for _, location := range artist.ConcertLocations {
			if strings.Contains(strings.ToLower(location), searchTerm) {
				filteredArtists = append(filteredArtists, artist)
				break
			}
		}
	}

	return filteredArtists
}

func FilterArtistsByCreationDates(artists []Artist, before1980, date1980to1990, date1990to2000, date2000to2010, after2010 bool) []Artist {
    var filteredArtists []Artist

    for _, artist := range artists {
        creationDate := artist.CreationDate

        // Apply the creation dates filter
        if (before1980 && creationDate <= 1980) ||
           (date1980to1990 && creationDate >= 1980 && creationDate <= 1990) ||
           (date1990to2000 && creationDate >= 1990 && creationDate <= 2000) ||
           (date2000to2010 && creationDate >= 2000 && creationDate <= 2010) ||
           (after2010 && creationDate > 2010) {
            filteredArtists = append(filteredArtists, artist)
        }
    }

    return filteredArtists
}