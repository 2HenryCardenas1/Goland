package main

import "fmt"

//Creacion de un arreglo
var tabla [10]int

func main() {
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)

	for _, valor := range tabla {
		fmt.Println(valor)
	}
}
