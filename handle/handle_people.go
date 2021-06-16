package handle

import (
	"encoding/json"
	"fmt"
	"go_api/defines"
	"io/ioutil"
	"log"
	"net/http"
)

func GetPeople(writer http.ResponseWriter, request *http.Request) {
	response, err := http.Get(defines.URL + "people")

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request people")
	}
	response_body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		log.Print("error parsing body")
	}

	var people defines.All_people

	if err := json.Unmarshal(response_body, &people); err != nil {
		fmt.Println("error parsing json", err)
	}

	for _, pers := range people.People {
		pers.GetWorld()
		fmt.Fprint(writer, pers)
	}
}
