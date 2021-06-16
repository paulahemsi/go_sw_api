package defines

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const URL = "https://swapi.dev/api/"

type Planet struct {
	Name       string `json:"name"`
	Terrain    string `json:"terrain"`
	Population string `json:"population"`
}

type Person struct {
	Name           string `json:"name"`
	Home_world_url string `json:"homeworld"`
	Home_world     Planet
}

type Specie struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Group    string `json:"classification"`
}

type Film struct {
	Name     string `json:"title"`
	Release  string `json:"release_date"`
	Episode  int    `json:"episode_id"`
	Director string `json:"director"`
}

func (people *Person) GetWorld() {
	res, err := http.Get(people.Home_world_url)
	if err != nil {
		log.Print("Failed to request world", err)
	}
	var res_body []byte
	if res_body, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("error parsing body", err)
	}
	json.Unmarshal(res_body, &people.Home_world)
}

type All_species struct {
	Specie []Specie `json:"results"`
}

type All_people struct {
	People []Person `json:"results"`
}

type All_films struct {
	Film []Film `json:"results"`
}
