package main

import (
	"os"
	"groupie_tracker/locations"
)

func main() {
	argument := os.Args
	url := argument[1]
	locations.OpenLocation(url)
}