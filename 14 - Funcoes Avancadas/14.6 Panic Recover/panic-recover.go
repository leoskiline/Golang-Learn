package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	} else if media < 6 {
		panic("A média é inferior a 6!")
	}
	return false
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 8))
	fmt.Println(alunoEstaAprovado(6, 5))
}
