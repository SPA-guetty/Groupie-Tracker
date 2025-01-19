package main

import (
	"fmt"
	"log"
	"net/http"
	"groupie_tracker/handlerfolder"
)

var port = ":8080"

func main() {
	// Server routes
	http.HandleFunc("/", handlerfolder.ArtHandler)
	
	// Static file server for images and css files
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
