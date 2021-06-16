package handle

import (
	"encoding/json"
	"fmt"
	"go_api/defines"
	"io/ioutil"
	"log"
	"net/http"
)

func GetFilms(writer http.ResponseWriter, request *http.Request) {
	response, err := http.Get(defines.URL + "films")

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request films")
	}
	response_body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Print("error parsing body")
	}

	var films defines.All_films

	if err := json.Unmarshal(response_body, &films); err != nil {
		fmt.Println("error parsing json", err)
	}

	fmt.Fprint(writer, films)
}
