package main //nombre de carpeta donde esta guardado

import (
	"fmt"
	"reflect"
)

func main() {

	//Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaración variables enteras
	base := 12          //:= indica que se crea y asigna
	var altura int = 14 //creamos variable y asignamos valor
	var area int        // creamos variable

	fmt.Println("Base:", base, "Altura:", altura, "Area:", area)

	//Zero values (valor por defecto para cada tipo)
	var a int     // 0
	var b float64 // 0
	var c string  // " "
	var d bool    // false

	fmt.Println(a, b, c, d)

	//Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado // infiere tipo

	fmt.Println("Area cuadrado:", areaCuadrado)
	fmt.Println(reflect.TypeOf(areaCuadrado))

}
