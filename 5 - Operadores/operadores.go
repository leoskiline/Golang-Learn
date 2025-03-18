package main

import "fmt"

func main() {
	// Aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 20
	var soma2 int16 = numero1 + numero2
	fmt.Println(soma2)

	// Fim dos Aritmeticos

	// Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	//Fim dos Operadores de Atribuição

	// Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// Fim dos Operadores Relacionais.

	// Operadores Logicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(true || false)
	fmt.Println(!false)
	// Fim dos Operadores Logicos

	// Operadores Unários
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 15
	fmt.Println(numero)

	numero *= 2
	numero /= 2
	numero %= 2
	fmt.Println(numero)
	// Fim dos Operadores Unários

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
