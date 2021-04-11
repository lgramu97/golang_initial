package main

import (
	"fmt"
	pk "go/src/mypackage"
)

func main() {
	// Importar paquete. Cuando una funcion comienza con Mayuscula, PUEDE EXPORTARSE Y UTILIZARSE EN OTROS PAQUETES.
	var mycar = pk.CarPublic
	mycar.Marca = "Ford"
	fmt.Println(mycar)

	var myCar = pk.carPrivate
	myCar.marca = "Ford"
	fmt.Println(myCar)

	pk.PrintMessage("HOLA MUNDO")

	pk.printMessage("HOLA MUNDO PRIVTE")

}
