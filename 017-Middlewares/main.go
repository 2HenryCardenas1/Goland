package main

import "fmt"

/*
Los middlewares nos permiten encapsular la ejecucion de una funcion pre existente con una funcion nueva
*/
var resultado int

func main() {
	fmt.Println("Inicio")
	resultado = operacionesMiddleware(sumar)(1, 3)
	fmt.Println(resultado)
	resultado = operacionesMiddleware(restar)(8, 3)
	fmt.Println(resultado)
	resultado = operacionesMiddleware(multiplicar)(5, 3)
	fmt.Println(resultado)
}

func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

//Los middleware tienen que ingresar y devolver los mismos tipos de datos de las funciones
func operacionesMiddleware(funcion func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion: ")
		return funcion(a, b)
	}
}
