package main

import "fmt"

func diaDaSemana(numero int) string {
	var dia string
	switch numero {
	case 1:
		dia = "Domingo"
	case 2:
		dia = "Segunda-Feira"
	case 3:
		dia = "Terça-Feira"
	case 4:
		dia = "Quarta-Feira"
	case 5:
		dia = "Quinta-Feira"
	case 6:
		dia = "Sexta-Feira"
	case 7:
		dia = "Sabado"
	default:
		dia = "Número Inválido"
	}
	return dia
}

func diaDaSemana2(numero int) string {
	var dia string
	switch {
	case numero == 1:
		dia = "Domingo"
		// fallthrough // Cai dentro da proxima condição
	case numero == 2:
		dia = "Segunda-Feira"
	case numero == 3:
		dia = "Terça-Feira"
	case numero == 4:
		dia = "Quarta-Feira"
	case numero == 5:
		dia = "Quinta-Feira"
	case numero == 6:
		dia = "Sexta-Feira"
	case numero == 7:
		dia = "Sabado"
	default:
		dia = "Número Inválido"
	}
	return dia
}

func main() {
	fmt.Println(diaDaSemana(200))
	fmt.Println(diaDaSemana2(5))
}
