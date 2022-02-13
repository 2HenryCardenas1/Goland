package main

import "fmt"

//Los mapas hace referencia a datos que se puden serializar
func main() {
	//Creacion de un mapa
	//make(map[tipo de dato que sera la clave]tipo de dato que sera el valor)
	paises := make(map[string]string)

	paises["Mexico"] = "Distrito-Federal"
	paises["Colombia"] = "Bogotá"
	fmt.Println(paises)

	//Creacion de otra forma un mapa (un mapa es como un diccionario en python)

	campeonato := map[string]int{
		"Real Madrid": 15,
		"Barcelona":   14,
		"Chealse":     29}
	fmt.Println(campeonato)

	//Adicionar elementos al mapa

	campeonato["River Plate"] = 25
	fmt.Println(campeonato)

	//Cambiando valores

	campeonato["Barcelona"] = 20
	fmt.Println(campeonato)

	//Eliminando valores de un mapa

	delete(campeonato, "Chealse")
	fmt.Println(campeonato)

	//Recorrer mapa

	for equipo, puntaje := range campeonato {
		fmt.Println(equipo, puntaje)
	}

	//Buscar si existe elemento en el mapa

	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe?= %t\n", puntaje, existe)
}
