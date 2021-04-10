package main //nombre de carpeta donde esta guardado

import "fmt"

func main() {

	// Declaración de variables
	helloMessage := "hello"
	worldMessage := "world"

	// Println (print con salto de linea)
	fmt.Println(helloMessage, worldMessage)

	// Printf (agrega funcion extra)
	nombre := "Tu vieja"
	curso := 500
	fmt.Printf("%s tiene mas de %d cursos  \n", nombre, curso)
	fmt.Printf("%v tiene mas de %v cursos  \n", nombre, curso) // Cuando no se sabe el tipo se pone (v)

	//Sprint (genera string pero no lo imprime, lo guarda)
	message := fmt.Sprintf("%s tiene mas de %d cursos ", nombre, curso) // Todo lo generado en Sprintf se guarda en message.
	fmt.Println(message)

	// Tipo de dato (con %T)
	fmt.Printf("helloMessage: %T \n", helloMessage)

	// Error message
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id) // %q es para string con " "
	fmt.Println(err.Error())                                 // fmt.Errorf retorna un string formateado que puede invocarse .Error()

	//Scan
	var input string
	fmt.Scan(&input) // Input por pantalla
	fmt.Printf("The string input contains: %s \n", input)

	// Sscan
	// Declaring two variables
	var nameSscan string
	var nameScan2 string
	var alphabet_count int
	// Calling the Sscan() function which
	// returns the number of elements
	// successfully scanned and error if
	// it persists
	nSscan, errSscan := fmt.Sscan("input 54 entendes?", &nameSscan, &alphabet_count, &nameScan2)
	// Below statements get executed if there is any error
	if errSscan != nil {
		panic(errSscan)
	}

	// Printing the number of elements and each elements also
	fmt.Printf("%d: %s, %d\n", nSscan, nameSscan, alphabet_count)

	// Sscanf
	// scans the argument string, storing successive space-separated values into successive
	// arguments as determined by the format.
	var name2 string
	var age int
	n, err := fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name2, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d: %s, %d\n", n, name2, age)

	// For more information: https://golang.org/pkg/fmt/
}
