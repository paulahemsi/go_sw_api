package main

import (
	"fmt"
	"go_api/handle"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/people", handle.GetPeople)
	http.HandleFunc("/films", handle.GetFilms)
	http.HandleFunc("/species", handle.GetSpecies)
	fmt.Println("Serving on 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
