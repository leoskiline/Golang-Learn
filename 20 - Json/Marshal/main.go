package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raça"`
	Idade uint   `json:"idade"`
}

func main() {
	cach := cachorro{
		Nome:  "Groot",
		Raca:  "Vira Lata",
		Idade: 5,
	}
	cachorroEmJSON, err := json.Marshal(cach)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorroEmJSON)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	fmt.Println(json_encode(c2))
}

func json_encode(v any) *bytes.Buffer {
	JSON, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return bytes.NewBuffer(JSON)
}
