package main

import (
	"fmt"
	"time"
)

//Con los canales podemos enviar informacion de una gorutine a otra funcion
// para que cada ejecucion paralela puedan interactuar con otra

func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1) //Con go hacemos llamadas asincronas: realizar varias tareas (hilos)
	fmt.Println("Llegue hasta acá")

	msg := <-canal1 //Hasta que no se reciba un valor en el canal1 no se para el codigo
	fmt.Println("Encontre un valor y me detuve", msg)

}

//Creamos la cuminacion entre el main y nuestra funcion bucle
func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 10000000000; i++ {

	}

	final := time.Now()
	//Sub es la funcion que retorna la duracion
	canal1 <- final.Sub(inicio) //<- con esto le asiganmos el valor a una variable chan
}
