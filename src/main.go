package main //nombre de carpeta donde esta guardado

import "fmt"

func main() {

	// Valores primitivos

	// Numeros enteros.
	var i1 int = 64   // depende del SO.
	var i2 int8 = 127 // -128 a 127
	var i3 int16 = 16 // 16 bits
	var i4 int32 = 32 // 32 bits
	var i5 int64 = 64 // 64 bits
	fmt.Println(i1, i2, i3, i4, i5)

	// Numeros enteros positivos. Valores mucho mas grandes.
	var u1 int = 64   // depende del SO.
	var u2 int8 = 8   // 0 a 2 a la 8 -1.
	var u3 int16 = 16 // 0 a 16 bits
	var u4 int32 = 32 // 0 a 32 bits
	var u5 int64 = 64 // 0 a 64 bits
	fmt.Println(u1, u2, u3, u4, u5)

	// Numeros decimales. Tiene mas alcance la parte decimal
	var f1 float32 = 32
	var f2 float64 = 64
	fmt.Println(f1, f2)

	// Text y booleanos
	var string string = "texto" // Entre " "
	var booleano bool = true    // true / false
	fmt.Println(string, booleano)

	// Numeros complejos
	var complex1 complex64 = 0  // Real e imaginario float 32
	var complex2 complex128 = 0 // real e imaginario float64
	fmt.Println(complex1, complex2)

}
