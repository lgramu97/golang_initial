package main

import "fmt"

//Channels are the pipes that connect concurrent goroutines.
//You can send values into channels from one goroutine and receive those values into another goroutine.
// Menos eficientes que go rutines pero mas eficientes en tiempo de desarrollo.

func say(text string, c chan<- string) { // del lado derecho <- indica que es de salida.
	c <- text // <- indica que se ingresa un dato
}

// Funcion que recibe un canal para enviar un mensaje y el mensaje a enviar.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func message(text string, c chan string) {
	c <- text
}

// Funcion que acepta un canal con un mensaje (pings), y un canal de salida (pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	channel := make(chan string, 1) // buena practica indicar el limite

	fmt.Println("Hello")

	go say("Bye", channel)

	fmt.Println(<-channel) // <- obtengo la salida (respuesta del segundo go rutine)
	// <- del lado derecho indica entrada, del lado izquierdo de salida

	//*********************** CHANNELS***************************//

	//Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	messages := make(chan string)

	//Send a value into a channel using the channel <- syntax.
	//Here we send "ping" to the messages channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	//	 The <-channel syntax receives a value from the channel.
	//	Here we’ll receive the "ping" message we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)

	pings := make(chan string, 1) // canal de entrada
	pongs := make(chan string, 1) // canal de salida
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs) // obtengo valor del canal pongs.

	//************** len y cap ***********//
	fmt.Println("*********** LEN Y CAP ***********")
	c := make(chan string, 2)
	c <- "Mensaje 1"
	//c <- "Mensaje 2"
	fmt.Println(len(c), cap(c)) // len: datos dentro del channel, cap: capacidad del channel
	c <- "Mensaje 2"

	// Range y Close
	close(c) // Indica que el canal se cierra, no va a recibir ningun otro dato. BUENA PRACTICA.
	//c <- "mensaje 3" //runtime error porque el canal esta cerrado.

	for message := range c { // Iterar por los valores de un canal
		fmt.Println(message)
	}

	// Select (permite esperar por multiples channels )
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Mensaje 1", email1)
	go message("Mensaje 2", email2)
	//No se cual responde primero, por eso utilizo select.
	// Itero por todos los canales
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email 1:", m1)

		case m2 := <-email2:
			{
				fmt.Println("Email recibido de email 2:", m2)
			}
		}
	}

}

//When we run the program the "ping" message is successfully passed from
//one goroutine to another via our channel.
// By default sends and receives block until both the sender and receiver are ready.
//This property allowed us to wait at the end of our program for the
//"ping" message without having to use any other synchronization.
