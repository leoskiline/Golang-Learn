package main

import "fmt"

func generica(interf any) {
	fmt.Println(interf)
}

func main() {
	generica("aaa")
	generica(12)
	generica(3.14)
	generica(true)
}
