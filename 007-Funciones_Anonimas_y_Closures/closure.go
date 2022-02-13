package main

import "fmt"

func closure() {
	var numero int
	fmt.Println("Ingrese la tabla que desea ver: ")
	fmt.Scanln(&numero)

	miTabla := Tabla(numero)
	for i := 1; i < 11; i++ {
		fmt.Printf("%d x %d = %d\n", numero, i, miTabla())
	}
}

//Esto es una funcion closure que retorna otra funcion
//Se

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia++
		return numero * secuencia
	}
}
