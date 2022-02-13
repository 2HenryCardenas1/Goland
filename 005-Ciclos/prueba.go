package main

import "fmt"

func prueba() {
	condicion := true
	var num int
	for condicion {
		fmt.Println("Estoy en un bucle infinito, en la iteracion: ", num)
		if num == 1000 {
			fmt.Println("Sali del bucle infinito llegue a la iteracion : ", num)
			break
		}
		num++
	}
}
