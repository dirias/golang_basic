package main

import (
	"fmt"
	"math"
)

func main() {
	// Calcular area de un cuadrado
	const base_cuadrado = 10
	areaCuadrado := base_cuadrado * base_cuadrado
	fmt.Println("Area del Cuadrado es:", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma", result)

	//Resta
	result = y - x
	fmt.Println("Resta", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplication", result)

	//Division
	result = y / x
	fmt.Println("Division", result)

	//Modulo
	result = y % x
	fmt.Println("Residuo", result)

	//Incremental
	x++
	fmt.Println("Incremental", x)

	//Decrecion
	x--
	fmt.Println("Decremental", x)

	//Area rectangulo
	// a = b x h
	base_rectangulo := 10
	altura_rectangulo := 15
	area_rectangulo := base_rectangulo * altura_rectangulo
	fmt.Println("Area rectangulo:", area_rectangulo)

	//Area trapecio
	// a = ((b+b)/2)*h
	base_mayor := 30
	base_menor := 20
	altura_trapecio := 10
	area_trapecio := ((base_mayor + base_menor) / 2) * altura_trapecio
	fmt.Println("Area trapecio:", area_trapecio)

	//Area circulo
	// a = pi x r^2
	pi := math.Pi
	radio := 20.0
	var area_circulo float64 = pi * math.Pow(2, radio)
	fmt.Println("Area circulo", area_circulo)
}
