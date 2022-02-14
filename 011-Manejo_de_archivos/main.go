package main

import (
	"bufio"
	"fmt" //libreria para leer entradas
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	fmt.Println("----")
	leoArchivo2()
	fmt.Println("----")
	crearArchivo()
	fmt.Println("----")
	editarArchivo()
}

//Con esta libreria no se puede editar ni manejar el archivo
func leoArchivo() {
	archivo, err := ioutil.ReadFile("./archivo.txt") //la funcio readFile necesita dos variables el archivo y el error

	if err != nil {
		fmt.Println("hubo un error")
	} else {
		fmt.Println(string(archivo))
	}

}

func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt")

	if err != nil {
		fmt.Println("Ocurrio un erros: ", err)
	} else {
		teclado := bufio.NewScanner(archivo)
		for teclado.Scan() { //Este for lo que hace es que si existe el archivo lo leera y guardara en memoria temporal
			registro := teclado.Text() //Aca se conbierte el texto caputaro en una cdena de texto
			fmt.Printf("Linea > %s\n", registro)
		}
	}
	archivo.Close() //Cerramos el archivo
}

func crearArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("hubo un error", err)

	} else {
		fmt.Fprintln(archivo, "Este es el texto nuevo ") //Este print lo que hace es grabar en el archivo
		archivo.Close()
		fmt.Printf("Archivo %s creado\n", archivo.Name())
	}

}

//Adicionar texto a un archivo
func editarArchivo() {
	fileName := "./nuevoArchivo.txt"
	if AdicionarTexto(fileName, "\n Esta es una nueva linea") == false {
		fmt.Println("Hubo un error al editar el texto")
	}
}

func AdicionarTexto(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644) //abrimos el archivo con permisos de escritura y lectura

	if err != nil {
		fmt.Println("Hubo un error al abrir el archivo ", err)
		return false
	}

	_, err = arch.WriteString(texto) //Este metodo es el que escribe en el archivo
	if err != nil {
		fmt.Println("Hubo un error ", err)
		return false
	}
	fmt.Println("Archivo editado con exito")
	arch.Close()
	return true
}
