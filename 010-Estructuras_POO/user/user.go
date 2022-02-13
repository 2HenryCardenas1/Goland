package user

import "time"

//Creando una estruturas(clases)
type Usuario struct {
	Id         int
	Nombre     string
	Fecha_Alta time.Time
	Status     bool
}

//	Creacion del constructor
// this *NombreClase => para apuntar y hacer referencia a la clase
func (this *Usuario) AltaUsuario(id int, nombre string, fecha_alta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.Fecha_Alta = fecha_alta
	this.Status = status
}
