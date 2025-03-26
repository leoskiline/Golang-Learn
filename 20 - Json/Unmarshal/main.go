package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"-"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Groot","raça":"Vira Lata","idade":5}`

	var c cachorro

	err := json.Unmarshal([]byte(cachorroEmJSON), &c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)

	cachorro2EmJSON := `{"nome": "Toby","raca":"Poodle"}`
	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro2EmJSON)

	fmt.Println(json_decode(`{"nome":"leonardo","cursos": [{"nome":"Sistemas de Informação"}]}`))
}

func json_decode(v string) map[string]any {
	c := make(map[string]any)
	if err := json.Unmarshal([]byte(v), &c); err != nil {
		log.Fatal(err)
	}
	return c
}
