package main

import "fmt"

func isPair(number int) bool {
	return (number % 2) == 0
}

func parOimpar(number int) {
	if isPair(number) {
		fmt.Printf("El numero %d es un numero par \n", number)
	} else {
		fmt.Printf("El numero %d es un numero impar \n", number)
	}
}

func validateUser(userName, password string) bool {
	return (userName == password)
}

func main() {

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Printf("Se cumple la condicion el valor es : %d \n", valor1)
	} else {
		fmt.Printf("Se cumple la condicion el valor es : %d \n", valor2)
	}

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := -2; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	numeroCheck := 15
	parOimpar(numeroCheck)
}
