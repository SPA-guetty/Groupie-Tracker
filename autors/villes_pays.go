package autors

import "fmt"

func Disband_Location(location string) (string, string) {
	var city string
	var country string
	writecity := true
	for _, chr := range location {
		if chr == '-' {
			writecity = false
		} else {
			if writecity {
				city = city + string(chr)
			} else {
				country = country + string(chr)
			}
		}
	}
	return city, country
}

func Contains(tab []string, str string) bool {
	for _, e := range tab {
		if e == str {
			return true
		}
	}
	return false
}

func Get_All_Locations() ([]string, []string) {
	var cities []string
	var countries []string
	locations := OpenAllLocations() 
	for _, location := range locations {
		for _, place := range location.Locations {
			city, country := Disband_Location(place)
			if !Contains(cities, city) {
				cities = append(cities, city)
			}
			if !Contains(countries, country) {
				countries = append(countries, country)
			}
		}
	} 
	return cities, countries
}

func Filter_By_Alp(tab []string) []string {
	var new_tab []string
	for range tab {
		firstval := tab[0]
		firstint := 0
		for i, e := range tab {
			if e < firstval {
				firstval = e
				firstint = i
			}
		}
		new_tab = append(new_tab, firstval)
		tab = append(tab[:firstint], tab[(firstint+1):]...)
	}
	return new_tab
}

func Print_Locations() {
	cities, countries := Get_All_Locations()
	cities = Filter_By_Alp(cities)
	countries = Filter_By_Alp(countries)
	for _, e := range cities {
		fmt.Println(e)
	}
	for _, e := range countries {
		fmt.Println(e)
	}
	fmt.Println("Nombre de villes:", len(cities))
	fmt.Println("Nombre de pays:", len(countries))
}