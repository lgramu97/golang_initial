package main

import "fmt"

func main() {

	// for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for while (for hasta que se cumpla una condicion)
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// for forever (ciclo infinito)
	counterForEver := 0
	for {
		fmt.Println("COunter forever", counterForEver)
		counterForEver++
		if counterForEver == 1000 {
			break
		}
	}

	// for descendente
	for i := counterForEver; i > 0; i-- {
		fmt.Println("Decremento", i)
	}

	// for range.
	//Here we use range to sum the numbers in a slice. Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//range on arrays and slices provides both the index and value for each entry.
	//Above we didn’t need the index, so we ignored it with the blank identifier _.
	//Sometimes we actually want the indexes though.
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
	//The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
