package main

import (
	"fmt"
	"sync"
	"time"
)

// A goroutine is a lightweight thread of execution.

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done() // Libero la go rutine
	fmt.Println(text)
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Como ejecutamos de forma concurrente?? Agregamos go

	var wg sync.WaitGroup // Acumula go rutienes y las libera de a una.
	go say("hello", &wg)

	wg.Add(2) // Agrego una go rutine al WaitGroup
	go say("world", &wg)

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	wg.Wait() // Espere hasta que terminen todos.
	//time.Sleep(time.Second * 1) // no es eficiente que para ejecutar la rutina concurrente hacer un wait.

	//************************************ GORUTINES ***********************************//
	//Suppose we have a function call f(s). Here’s how we’d call that in the usual way,
	//running it synchronously.

	f("direct")

	//To invoke this function in a goroutine, use go f(s).
	//This new goroutine will execute concurrently with the calling one.
	go f("goroutine")

	//You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//Our two function calls are running asynchronously in separate goroutines now.
	//Wait for them to finish (for a more robust approach, use a WaitGroup).

	time.Sleep(time.Second)
	fmt.Println("done")
}

//When we run this program, we see the output of the blocking call first,
//then the interleaved output of the two goroutines.
//This interleaving reflects the goroutines being run concurrently by the Go runtime.
