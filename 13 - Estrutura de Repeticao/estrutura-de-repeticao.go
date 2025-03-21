package main

import (
	"fmt"
)

func main() {
	// i := 0
	// for i < 100 {
	// 	time.Sleep(time.Second / 10)
	// 	i += 4
	// 	fmt.Println(i)
	// }

	// for j := 1; j <= 10; j++ {
	// 	time.Sleep(time.Second / 5)
	// 	fmt.Println(j)
	// }

	nomes := [3]string{"Leo", "Davi", "Jose"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, letra := range "PALAVRA" {
		fmt.Print(string(letra))
	}
	fmt.Println()

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	type usuarioStruct struct {
		nome      string
		sobrenome string
	}

	usuario2 := usuarioStruct{"Zé", "Junior"}
	fmt.Println(usuario2)
}
