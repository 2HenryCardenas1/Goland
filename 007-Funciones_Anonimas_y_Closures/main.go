package main

import "fmt"

//Las funciones anonimas son funciones que no tienen nombre,
//se usan para ocultar codigo dentro de una variable y
// en estas funciones se pueden modificar en diferentes partes del codigo

//Ejemplo de variable anonima
var Calculo func(int, int) int = func(num int, num2 int) int {
	return num + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5, 7))

	//Resta
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Resta 10-5 = %d\n", Calculo(10, 5))

	//Division
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}

	fmt.Printf("Dividido 10/5 = %d\n", Calculo(10, 5))

	//Multiplicacion
	Calculo = func(num1 int, num2 int) int {
		return num1 * num2
	}

	fmt.Printf("Multiplico 10*5 = %d\n", Calculo(10, 5))

	Operaciones()
}

//2do ejemplo
func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}
