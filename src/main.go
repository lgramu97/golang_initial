package main //nombre de carpeta donde esta guardado

import "fmt"

// Declaración de función
func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) { // a y b son int.
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func returnDoubleValue(a, b float32) (c, d float32) {
	return a * b, a / b
}

func getAreaRectangulo(base, altura float32) float32 {
	return base * altura
}

func getAreaRectanguloAndMetrics(base, altura float32) (a, b, result float32) {
	return base, altura, getAreaRectangulo(base, altura)
}

func main() {
	normalFunction("Hello World1")
	tripleArgument(1, 2, "Hola")
	result := returnValue(4)
	fmt.Println(result)

	value1, value2 := returnDoubleValue(2, 4)
	fmt.Printf("Value 1: %f , Value 2: %f \n", value1, value2)

	// Si no me intersa un valor
	value3, _ := returnDoubleValue(2, 4)
	fmt.Printf("Value 1: %f  \n", value3)

	fmt.Printf("Area de un rectangulo: %f \n", getAreaRectangulo(2, 3))

	base, altura, area := getAreaRectanguloAndMetrics(4, 4)
	fmt.Printf("El Area de un rectangulo de Base %.2f y Altura %.2f es : %.4f \n", base, altura, area)

}
