package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Joao", "Pedro", 20, 1.75}
	fmt.Println(p1)
	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1.nome, e1.altura)
}
