package main

import "fmt"

func main() {
	fmt.Println("Working with arrays\n")
	var array [5] int

	array[0] = 1
	array[1] = 2

	fmt.Println(array, cap(array), len(array))

	//slice
	fmt.Println("\nWorking with slices\n")
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice,cap(slice), len(slice))
	//El primer numero es inclusivo y el segundo exclusivo
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Los arrays son inmutables, sin embargo, se le pueden modificar los elementos ya existente.
	//Los slices no mutables, se le pueden agregar o quitar elementos

	//append - agregar un elemento nuevo
	slice = append(slice, 7)
	fmt.Println(slice,cap(slice), len(slice))

	//append - agregar un otro slice
	nuevoSlice := [] int {9, 10, 11}
	slice = append(slice, nuevoSlice...)
	fmt.Println(slice,cap(slice), len(slice))

}