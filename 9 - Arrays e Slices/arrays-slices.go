package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]string
	array1[0] = "Pos 1"
	fmt.Println(array1)
	array2 := [5]string{"Pos 1", "Pos 2", "Pos 3", "Pos 4", "Pos 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	slice = append(slice, 5)
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)
}
