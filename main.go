package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Commune struct {
	Nom        string `json:"nom"`
	Code       string `json:"code"`
	Population int    `json:"pop"`
}

func main() {
	const API_URL string = "https://geo.api.gouv.fr"
	URI_Communes := "/communes?codePostal=12000"

	res, err := http.Get(API_URL + URI_Communes)
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
		log.Printf("Commune : %v\n Code : %v\n Population : %v\n",
			commune.Nom, commune.Code, commune.Population)
	}

}
