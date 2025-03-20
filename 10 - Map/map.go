package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Leonardo",
			"ultimo":   "Custodio",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "curso")
	fmt.Println(usuario2)
	usuario2["signo"] = map[string]string{
		"nome": "Touro",
		"mes":  "Abril",
	}
	fmt.Println(usuario2)
}
