package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) bool {
	// version 1 comparando 1 a 1
	cumple := true
	text = strings.ToLower(text)
	for i := range text {
		if text[i] == text[len(text)-i-1] {
			continue
		} else {
			cumple = false
			break
		}
	}
	return cumple
}

func isPalindromo2(text string) bool {
	// version 2 contruyendo un nuevo string y luego comparando
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	return strings.ToLower(textReverse) == strings.ToLower(text)
}

func main() {
	// range iterates over elements in a variety of data structures.
	//Let’s see how to use range with some of the data structures we’ve already learned.

	//Here we use range to sum the numbers in a slice. Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//range on arrays and slices provides both the index and value for each entry.
	// Above we didn’t need the index, so we ignored it with the blank identifier _.
	// Sometimes we actually want the indexes though.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	//*************************************************Platzi**********************************//
	slice := []string{"hola", "mundo", "chau"}

	// Las 2 variables
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	// Solo la palabra
	for _, valor := range slice {
		fmt.Println(valor)
	}

	// Solo el indice
	for i := range slice {
		fmt.Println(i)
	}

	palindromo := "neuqueN"
	fmt.Printf("Es un palindromo la palabra %s: %t \n", palindromo, isPalindromo(palindromo))
	fmt.Printf("Es un palindromo la palabra %s: %t \n", palindromo, isPalindromo2(palindromo))

}
