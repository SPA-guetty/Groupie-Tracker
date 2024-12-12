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