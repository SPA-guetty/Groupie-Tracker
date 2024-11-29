package main

import (
	"groupie_tracker/relations"
	"os"
)

func main() {
	args := os.Args
	id := args[1]
	relations.OpenRelation(id)
}