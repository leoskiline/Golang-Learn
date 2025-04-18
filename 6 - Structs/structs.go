package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "Leonardo"
	u.idade = 28
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos girassois", 50}
	usuario2 := usuario{"Jose", 50, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Joseph"}
	fmt.Println(usuario3)
}
