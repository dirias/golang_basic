package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripeArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("value1 y value2:", value1, value2)

	//si solo queremos retornar un valor cuando usamos doubleReturn, podemos usar _

	//value1, _ := doubleReturn(2)
	//fmt.Println("value1:", value1)
}