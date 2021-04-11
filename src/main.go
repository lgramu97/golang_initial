package main

import "fmt"

func main() {
	// ARRAY INMUTABLE, SLICE MUTABLE

	// Array
	var array [4]int // inicializar en 0 arreglo
	array[0] = 1
	array[1] = 2
	fmt.Println(array, "Longitud:", len(array), "Maxima capacidad:", cap(array))

	// Slice (arreglo pero no se indica la capacidad, es una lista)
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, "Longitud:", len(slice), "Maxima capacidad:", cap(slice))

	// Metodos Slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])  // Numero despues de : es exlusivo
	fmt.Println(slice[2:4]) // inclusivo : exclusivo
	fmt.Println(slice[4:])
	fmt.Println(slice[:])

	// Append 1 elemento
	slice = append(slice, 7)
	fmt.Println(slice[:])

	// Append multiples elementos
	newSlice := []int{8, 9, 10, 11}
	slice = append(slice, newSlice...)
	fmt.Println(slice[:])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	//Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Copy slices

	//Slices can also be copy’d. Here we create an empty slice c of the same length as slice
	//and copy into c from s.
	c := make([]int, len(slice))
	copy(c, slice)
	fmt.Println("cpy of slice:", c)

}
