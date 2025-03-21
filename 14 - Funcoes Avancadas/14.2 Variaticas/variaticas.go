package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 1, 1, 123, 2, 1123, 123, 123, 123, 13, 1)
	fmt.Println(totalSoma)

	escrever("Olá Mundo", 1, 2, 3, 455, 131, 321)
}
