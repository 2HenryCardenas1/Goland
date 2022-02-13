package main

import "fmt"

func ejemplo2() {

	fmt.Println("ARRAY")
	tabla := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(tabla); i++ {
		fmt.Println(i)
	}

	fmt.Println("MATRIZ")

	var matriz [5][5]int

	fmt.Println(matriz)

	//Slice

	fmt.Println("Vector Dinamico")

	slice := []int{2, 5, 6}

	fmt.Println(slice)

	variante()
	variente2()

}

func variante() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:] //Creacion de vector desde la posicion 3 hasta la ultima de (elemntos)

	fmt.Println(porcion)
}

//make([]tipo de arreglo,cuantos elementos desea recervar,tamaño en memoria)
func variente2() {
	elemento := make([]int, 5, 20)
	fmt.Printf("Largo %d Capacidad %d", len(elemento), cap(elemento)) //cap funciona para ver la capacidad del arreglo
	nums := make([]int, 0, 0)
	for i := 1; i <= 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
