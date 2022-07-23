package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PokemonSpecies struct {
	Name string `json:"name"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

func main() {
	fmt.Println("Start API Consumption")

	URL := "http://pokeapi.co/api/v2/pokedex/kanto/"

	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
		data := Response{}
		jsonErr := json.Unmarshal(body, &data)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		fmt.Println("Name :", data.Name)
		for i := 0; i < len(data.Pokemon); i++ {
			fmt.Println(data.Pokemon[i].Species.Name)
		}
	}
}
