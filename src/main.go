package main

import "fmt"

type rect struct {
	width, height int
}

//This area method has a receiver type of *rect.
func (r *rect) area() int {
	return r.width * r.height
}

//Methods can be defined for either pointer or value receiver types.
//Here’s an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

type pc struct {
	ram   int
	disk  float32
	marca string
}

// Parseo de datos de un struct para imprimir "bonito". Overwrite metodo fmt.Println(pc) para struct pc.
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d RAM, un disco de %0.1f GB y mi pc es de la marca %s \n", myPc.ram, myPc.disk, myPc.marca)
}

// Definir funcion para un struct. func (name struct_name) name_function(){cuerpo}
func (myPc pc) ping() {
	fmt.Println(myPc.marca, "Pong")
}

// Si se pasa con el * es por referencia, se modifica la memoria directametne.
func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

//We’ll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr.
// zeroval has an int parameter, so arguments will be passed to it by value.
// zeroval will get a copy of ival distinct from the one in the calling function.
func zeroval(ival int) {
	ival = 0
}

//zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
// The *iptr code in the function body then dereferences the pointer from its memory
//address to the current value at that address.
//Assigning a value to a dereferenced pointer changes the value at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	// Punteros y Structs

	a := 50
	b := &a //la direccion de a

	fmt.Println(a, b)

	// Como acceder el valor?
	fmt.Println(*b) // Utilizo el * para saber el valor
	fmt.Println(&b) // utilizo el & para saber la direccion.

	*b = 100 // Cambio el valor de a
	fmt.Println(a)

	var myPc pc
	myPc.disk = 200
	myPc.marca = "HP"
	myPc.ram = 12
	fmt.Println(myPc)
	myPc.ping()

	// Estado actual
	fmt.Println(myPc)

	//Duplico ram con el metodo
	myPc.duplicateRam()
	fmt.Println(myPc)

	//Go supports pointers, allowing you to pass references to values and records within your program.
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	//The &i syntax gives the memory address of i, i.e. a pointer to i.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	//Pointers can be printed too.
	fmt.Println("pointer:", &i)

	//zeroval doesn’t change the i in main, but zeroptr does because it has a reference to the memory address
	//for that variable.

	// Personalizar output struct. Sobreescribo el metodo String()
	myNewPc := pc{ram: 12, disk: 45, marca: "HP"}
	fmt.Println(myNewPc)

	//************** METHODS-FUNCTIONS **************//
	fmt.Println("------------------Functions---------------")

	r := rect{width: 10, height: 5}

	//Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	//Go automatically handles conversion between values and pointers for method calls.
	//You may want to use a pointer receiver type to avoid copying on method calls or
	//to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

}
