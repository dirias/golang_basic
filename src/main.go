//https://pkg.go.dev/fmt
package main

import (
	"fmt"
)

func main() {
	helloMessage := "Hello"
	worldMessage := "world"

	//Println - agrega el salto de linea al final
	fmt.Println(helloMessage, worldMessage)

	nombre := "Didier"
	cursos := 400

	// Printf - %s string, %d un entero , %v undifiend
	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos \n", nombre, cursos)

	//Sprintf - Genera un string y lo guardo, no lo imprime en consola

	message := fmt.Sprintf("%s tiene mas de %d cursos \n", nombre, cursos)
	fmt.Printf(message)

	//Imprimir el tipo de datos
	
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
