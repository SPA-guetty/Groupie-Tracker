package autors

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
		tabart = Filter_By_Name(tabart)
		for i := range tabart {
			new_tab = append(new_tab, tabart[i])
			tab = append(tab[:(tabint[i]-i)], tab[(tabint[i]-i+1):]...)
		}
	}
	return new_tab
}