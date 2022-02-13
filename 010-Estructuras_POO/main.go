package main

import (
	"fmt"
	"time"

	usr "Goland/010-Estructuras_POO/user" //Llamando a clases o archivos
)

//Creando una estruturas(clases)
//type usuario struct {
//	Id         int
//	Nombre     string
//	Fecha_Alta time.Time
//	Status     bool
//}

//Herencia
//Aca creamos una nueva estructura y le indicamos que
//va a heredar de la estructura user creada en la carpera User
type persona struct {
	usr.Usuario
}

func main() {
	//	Llamado de estructuras(clases)
	//usuario := new(usuario)
	usuario := new(persona)
	usuario.AltaUsuario(1, "Pepito Pepon", time.Now(), true)
	fmt.Print(usuario.Usuario)
}
