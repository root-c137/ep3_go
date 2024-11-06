package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Commune struct {
	Name    string
	Country string
	State   string
	Lat     float32
	Lan     float32
}

func main() {
	CityName := "Lyon"
	const BASE_URL string = "http://api.openweathermap.org/geo/1.0/direct?"
	const API_KEY = "c732a4f732342956ec521490b59a7dce"
	URI := "q=" + CityName + "&limit=1&appid=" + API_KEY

	res, err := http.Get(BASE_URL + URI)
	if err != nil {
		log.Fatalf("La récuprération des données a échoué... : %v", err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("La récuprération des données a échoué... : \nStatus code : %v\nerreur : %v",
			res.StatusCode, err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("La récuprération des données a échoué... : %v", err)
	}

	var communes []Commune
	err = json.Unmarshal(body, &communes)
	if err != nil {
		log.Fatalf("La sérialisation n'a pas fonctionné... : %v", err)
	}

	for _, commune := range communes {
		log.Printf("\n\nCommune : %v\nCode : %v\nPopulation : %v\nLatitude : %v\nLongitude : %v\n\n",
			commune.Name, commune.Country, commune.State, commune.Lat, commune.Lan)
	}

}
