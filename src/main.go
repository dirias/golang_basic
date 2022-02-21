package main

import (
	"fmt"
)
//Es la manera de crear clases en go
type car struct{
	brand string
	year int
}

func main() {
	//Manera 1 de instanciar clases
	myCar := car{brand: "Ford", year: 2022}
	fmt.Println(myCar)
	//Manera 2 de instanciar clases
	var otherCar car //Lo instancia como un constuctor vacio
	otherCar.brand = "Ferrari" //Al no indicar el otro valor le asigna el nulo por defecto
	fmt.Println(otherCar)



}