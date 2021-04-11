package main

import "fmt"

func main() {
	// Diccionarios o Maps

	// Crear map
	edades := make(map[string]int)
	// Agregar datos
	edades["Jose"] = 14
	edades["Pepe"] = 20
	// Imprimir
	fmt.Println(edades)

	// Recorrer map
	for k, value := range edades {
		fmt.Println(k, value)
	}

	// MUCHO CUIDADO!!
	// Si la clave no existe, se muestra el valor del zero_value
	value, ok := edades["Josee"] // Si pongo 2 valores, el segundo valor retornado es si existe.
	fmt.Println(value, ok)

	value, ok = edades["Jose"]
	fmt.Println(value, ok)

}
