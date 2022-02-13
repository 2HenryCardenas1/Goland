package main

import (
	"bufio"
	"fmt"
	"os" //Libreria para el manejo de acciones del sistema
)

var numero1, numero2 int
var texto string

func main() {
	fmt.Println("Ingrese un numero: ")
	fmt.Scanln(&numero1) //La funcion Scanln se usa para capturar entradas por teclado
	//Para llamar a una variable se debe de llamar con &variable
	fmt.Println("Ingrese un numero: ")
	fmt.Scanln(&numero2)

	fmt.Println("Descripcion: ")

	teclado := bufio.NewScanner(os.Stdin) //Aca indicamos que el usuario puede escribir por teclado

	if teclado.Scan() { //Comprobamos que el usuario pulso alguna tecla
		texto = teclado.Text() //Le damos el valor ingresado a la variable texo

	}
	fmt.Println("Lo que ingreso el usuario fue: ", texto)
	fmt.Println("El resultado de la suma de los numeros es: ", numero1+numero2)
}
