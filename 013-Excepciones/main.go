package main

import (
	"log"
)

func main() {
	/*archivo := "archivo.txt"
	file, err := os.Open(archivo)

	defer file.Close() //antes de salir de la funcion esto cierra el archivo

	if err != nil {
		fmt.Println("Error abriendo el archivo")
		os.Exit(1)
	}
	fmt.Println("---")*/
	ejemploPanic()
}

func ejemploPanic() {

	defer func() {
		reco := recover() //recover verifica si hubo un panic en el codigo
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero Panic\n%v", reco) //%v se usa para llamar un valor variable(que cambia)

		}

	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1 ")
	}
}
