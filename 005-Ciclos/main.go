package main

import (
	"fmt"
)

func main() {
	var i = 1
	for i <= 10 {

		fmt.Println(i)
		i++
	}
	fmt.Println()
	for2()
	fmt.Println()
	whileGo()
	fmt.Println()
	continueGO()
	fmt.Println()
	gotoGo()
}

//For de la forma de java
func for2() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

//While no existe en GO

func whileGo() {
	var numero int
	for {
		fmt.Println("Ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			fmt.Println("El numero secreto era el 100")
			break
		}
		fmt.Println("Intente de nuevo")
	}
}

//Uso del continue,al encontrar el continue lo que hace es regresar al inicio del ciclo for

func continueGO() {
	var i = 0

	for i < 10 {
		fmt.Printf("\n Valor de i: %d\n", i)
		if i == 5 {
			fmt.Printf("Multiplicamos por 2\n")
			i = i * 2
			fmt.Printf("Upa encontramos el continue\nEl valor final de i es = %d", i)
			continue //Al encontrar el continue lo que hace es regresar al inicio del ciclo for
		}
		fmt.Printf("Ya paso el ciclo i = %d", i)
		i++
	}
}

//El goto nos envia a la parte del codigo especificado y sigue corriendo el codigo
func gotoGo() {
	var i = 0
RUTINA: //Al volver aca toma el valor de la variable donde llamamos el goto
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Printf("Voy a RUTINA sumando 2 a el valor de i es decir que el siguiente nuemro sera 6\n")
			goto RUTINA
		}
		fmt.Printf("Valor de i:%d\n", i)
		i++
	}

}
