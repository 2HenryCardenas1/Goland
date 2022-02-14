package main

import (
	"fmt"
	"strings"
	"time"
)

//Go routines son promesas
func main() {
	//Con go hacemos llamadas asincronas: realizar varias tareas (hilos)
	go miNombreLentoo("Pepito Perez")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Println("Escribe algo")
	fmt.Scanln(&x)
}
func miNombreLentoo(nombre string) {
	//Con esta libreria(strings.Split) logramos convertir un string en un vector
	//Separamos cada uno de sus caracteres
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond) //Espere un segundo por iteracion
		fmt.Println(letra)
	}
}
