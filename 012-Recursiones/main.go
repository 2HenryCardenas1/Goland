package main

import "fmt"

func main() {
	exponencia(2)
}

//Funcion recursiva

func exponencia(numero int) {
	if numero > 1000000 {
		return //Con return finalizamos la condicion
	}
	fmt.Println(numero)

	//recurcion

	exponencia(numero * 2)
}
