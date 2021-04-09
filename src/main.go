package main //nombre de carpeta donde esta guardado

import (
	"fmt"
	"math"
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

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = x - y // Sin los : ya que se definió en la primer asignación
	fmt.Println("Resta:", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiṕlicación:", result)

	// División
	result = y / x
	fmt.Println("División:", result)

	// Modulo (resto)
	result = y % x
	fmt.Println("Módulo:", result)

	// Incremental
	result += 1
	result++
	fmt.Println("Incremental:", result)

	// Decremental
	result -= 1
	result--
	fmt.Println("Decremento:", result)

	// Area de Rectangulo,
	baseRectangulo := 20
	alturaRectangulo := 10
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("Area de un rectangulo:", areaRectangulo)

	//Area Trapesio
	baseTrapesio := 10
	alturaTrapesio := 10
	areaTrapesio := (baseTrapesio * alturaTrapesio) / 2
	fmt.Println("Area de un trapesio:", areaTrapesio)

	//Area Circulo
	var radioCirculo float64 = 20
	areaCirculo := math.Pi * math.Pow(radioCirculo, 2)
	fmt.Println("Area circulo:", areaCirculo)

}
