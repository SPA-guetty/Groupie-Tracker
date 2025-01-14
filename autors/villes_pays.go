package autors

import "fmt"

type country struct {
	Country		string
	Cities		[]string
}

func Disband_Location(location string) (string, string) {
	var city string
	var country string
	writecity := true
	for _, chr := range location {
		if chr == '(' {
			writecity = false
			city = city[:(len(city)-1)]
		} else {
			if writecity {
				city = city + string(chr)
			} else {
				if string(chr) != ")" {
					country = country + string(chr)
				}
			}
		}
	}
	return city, country
}

func Contains_country(countries []country, nation string) bool {
	for _, e := range countries {
		if e.Country == nation {
			return true
		}
	}
	return false
}

func Contains_city(countries []country, nation string, city string) int {
	index := -1
	for i, e := range countries {
		if e.Country == nation {
			index = i
			for _, cities := range e.Cities {
				if cities == city {
					return -1
				}
			}
		}
	}
	return index
}

func Get_All_Locations() []country {
	var world []country
	locations := OpenAllLocations() 
	for _, location := range locations {
		for _, place := range location.Locations {
			city, nation := Disband_Location(place)
			if !Contains_country(world, nation) {
				var newcountry country
				newcountry.Country = nation
				world = append(world, newcountry)
			}
			index := Contains_city(world, nation, city)
			if index != -1 {
				world[index].Cities = append(world[index].Cities, city)
			}
		}
	} 
	return world
}

func Filter_City_By_Alp(tab []string) []string {
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

func Filter_By_Alp(tab []country) []country {
	var new_tab []country
	for range tab {
		firstval := tab[0]
		firstint := 0
		for i, e := range tab {
			if e.Country < firstval.Country {
				firstval = e
				firstint = i
			}
		}
		firstval.Cities = Filter_City_By_Alp(firstval.Cities)
		new_tab = append(new_tab, firstval)
		tab = append(tab[:firstint], tab[(firstint+1):]...)
	}
	return new_tab
}

func Find_Locations() {
	world := Get_All_Locations()
	world = Filter_By_Alp(world)/*
	countries = Filter_By_Alp(countries)*/
	for _, e := range world {
		fmt.Println(e)
	}
	/*for _, e := range countries {
		fmt.Println(e)
	}
	fmt.Println("Nombre de villes:", len(cities))
	fmt.Println("Nombre de pays:", len(countries))*/

}