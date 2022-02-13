package main

import "fmt"

func main() {
	fmt.Println(uno(5))
	numero, estado := dos(1)
	fmt.Println(numero, estado)
	fmt.Println("--------")
	fmt.Println(Calculo(5, 12))
	fmt.Println("--------")
	fmt.Println(Calculo(6, 12, 324, 54))
	fmt.Println("--------")
	fmt.Println(Calculo(32, 12, 22))
	fmt.Println("--------")
	fmt.Println(Calculo(8, 12))
	fmt.Println("--------")
	fmt.Println(Calculo(1))
}

func uno(numero int) int { // func -nombre de la funcion- -(parametros)- -tipo de valor a devolver-
	return numero * 2
}

//Tipos de datos a devolver
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}

}

//Parametros variables
// _ el guion bajo sirve para almacenar variables basura que solo se usan una vez o no se necesitan guardar
// range sirve para recorrer arreglos y devuelven dos valores
func Calculo(arreglo ...int) (txt string, total int) {
	total = 0
	for item, num := range arreglo { // Aca le decimos que me recorra los valores del arreglo y se los proporcione a item(es la posicion) y a num(valor encontrado)
		total = total + num
		fmt.Printf("esta es la posicion %d y tiene el valor de %d \n", item, num)

	}
	txt = "La suma entre los valores del arreglo es : "
	return txt, total
}
