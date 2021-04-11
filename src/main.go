package main

import "fmt"

func main() {
	// Defer (linea que se ejecuta ultimo)
	// utilizado para cerrar conexiones, archivos, base de datos.
	// por ejemplo defer closeFile()
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue // A pesar de un error quiero que continue por ej.
		}

		// Break
		if i >= 8 {
			fmt.Println("Muy alto el numero")
			break // Quiero que corte el for.
		}

	}
}
