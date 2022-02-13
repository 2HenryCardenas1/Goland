package main

import (
	"fmt"
	"strconv"
)

//Una variable publica para poder acceder a ella fuera de este paquete se inicializa asi
var Numero int

var texto string

//Asi se puede inicializar las variables
var status bool = true

// En go se puede crear y asignarle valores en una sola linea a variables
func main() {
	numero1, numero2, numero3, texto := 2, 5, 67, "Este es mi texto"

	var moneda float32 = 0
	//Asi se parcea los tipos de datos para pasarlos de uno a otro: ejemplo float32 a int
	numero2 = int(moneda)

	//pasando de numero a texto con fmt.SprintF()
	var moneda2 = 3

	texto = fmt.Sprintf("%d", moneda2)

	//pasando de entero a estring usando la libreria strconv
	texto2 := "Nuevo texto"
	var moneda3 int64 = 40

	texto2 = strconv.Itoa(int(moneda3))

	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println()
	fmt.Println(texto)
	fmt.Println()
	fmt.Println(texto2)

	//Asi se llaman los metodos a una funcion principal

	mostrarStatus()

}

func mostrarStatus() {
	fmt.Println(status)
}
