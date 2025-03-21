package main

import "fmt"

func main() {

	texto := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 10)
	}("Olá Mundo")

	fmt.Println(texto)
}
