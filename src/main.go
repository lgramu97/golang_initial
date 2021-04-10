package main

import (
	"fmt"
	"time"
)

func main() {

	// Definicion y parseo de variable
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Switch sin condicion (cuando quiero anidar multiples condiciones)
	value := 54
	switch {
	case value > 100:
		fmt.Println("Es mayor que 100")
	case value < 0:
		fmt.Println("Es menor que 0")
	default:
		fmt.Println("No condicion")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//switch without an expression is an alternate way to express if/else logic.
	//Here we also show how the case expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//A type switch compares types instead of values.
	//You can use this to discover the type of an interface value.
	//In this example, the variable t will have the type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
