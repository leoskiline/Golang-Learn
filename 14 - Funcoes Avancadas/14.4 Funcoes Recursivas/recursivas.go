package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	for posicao := uint(0); posicao < 10; posicao++ {
		fmt.Print(fibonacci(posicao), " ")
	}

}
