package main

import (
	"fmt"
)


func main() {

	m := make(map[string] int)
	m["Didier"] = 24
	m["Didiana"] = 18
	//Ocurre de forma concurrente, por ende no se va a hacer el output en orden
	for i, v := range m{
		fmt.Println(i, v)
	}
	//El ok retorna el true si existe y false si no, debido a que si se busca un valor que no existe
	//Se retorna 0
	value, ok := m["Didier"]
	fmt.Println(value, ok)

}