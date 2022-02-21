package main

import (
	"fmt"
	"strings"
)

func isPalyndrome(world string){
	var textReverse string
	
	for i := len(world) -1; i >= 0; i--{
		//si no se convierte a string directamente va a retornar el codigo ASCII
		textReverse += string(world[i])
	}
	if strings.ToLower(world) == strings.ToLower(textReverse){
		fmt.Println("Es palindromo")
	}else{
		fmt.Println("No es palindromo")

	}

}
func main() {

	world := "amor a romA"
	
	isPalyndrome(world)

}