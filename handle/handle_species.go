package handle

import (
	"encoding/json"
	"fmt"
	"go_api/defines"
	"io/ioutil"
	"log"
	"net/http"
)

func GetSpecies(writer http.ResponseWriter, request *http.Request) {
	response, err := http.Get(defines.URL + "species")

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request films")
	}
	response_body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Print("error parsing body")
	}

	var species defines.All_species

	if err := json.Unmarshal(response_body, &species); err != nil {
		fmt.Println("error parsing json", err)
	}

	fmt.Fprint(writer, species)
}
