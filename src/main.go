package main

import "fmt"


func main() {

	fmt.Println()
	fmt.Println("=================")
	fmt.Println("El condicional IF")
	fmt.Println("=================")
	fmt.Println()

	fmt.Println("Como se indica en los videos la instrucción IF se declara igual que en C, Java, kotlin pero sin añadir entre" +
		" paréntesis las condiciones. El caso es que se puede con estos paréntesis y funciona, compila, pero el corrector" +
		" de estilo del editor los quita.")
	fmt.Println()

	fmt.Println("En el siguiente ejemplo se usa el operador de comparación de igualdad y los operadores lógicos AND '&&' y OR '||'")
	fmt.Println()

	fmt.Println("Declaración:")
	fmt.Println()

	fmt.Println("if (valor1 == 1 && valor2 == 2) || valor3 == 3 { fmt.Println(\"IF > Se cumplen las condiciones.\") } else { fmt.Println(\"ELSE > No se cumplen las condiciones del IF\") }")
	fmt.Println()

	fmt.Println("Para la declaración de valores en el código, el resultado es:")
	fmt.Println()

	const valor1 int8 = 11
	var valor2 int16 = 22
	valor3 := 33

	if (valor1 == 1 && valor2 == 2) || valor3 == 3 {
		fmt.Println("IF > Se cumplen las condiciones.")
	} else {
		fmt.Println("ELSE > No se cumplen las condiciones del IF")
	}

	fmt.Println()

	fmt.Println("Finalmente dejo un experimento con FMT y decisiones lógicas")
	fmt.Println()

	fmt.Printf("Con el valor >>> %t <<< confirmo que la siguiente sentencia da Verdadero:\n" + 
	"(valor1 == 1 && valor2 == 2) || valor3 == 33\n\n", (valor1 == 1 && valor2 == 2) || valor3 == 33)
}