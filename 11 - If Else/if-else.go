package main

import "fmt"

func main() {

	numero := 10

	if numero > 15 {
		fmt.Println("É Maior de 15")
	} else {
		fmt.Println("É Menor de 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	}
}
