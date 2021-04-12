package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// Primero instalamos el paquete:  go get -v -u  github.com/labstack/echo

	// Creamos el modulo. go mod init github.com/lgramu97/golang_initial
	// Es buena practica poner como modulo el nombre de la carpeta donde se va a alojar.

	// Instanciamos echo
	e := echo.New()

	// Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
