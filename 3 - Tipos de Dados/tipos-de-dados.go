package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = 100
	fmt.Println(numero)

	var numero2 uint = 1000
	fmt.Println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	// INT8 = BYTE
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123
	fmt.Println(numeroReal2)

	var str string = "Texto"
	fmt.Println(str)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro genérico")
	fmt.Println(erro)

	var char byte = 'A'
	fmt.Println(char)

	var texto rune = 'A'
	fmt.Println(texto)

	var numero5 complex64 = 123456
	fmt.Println(numero5)

	var numero6 complex128 = 123456
	fmt.Println(numero6)
}
