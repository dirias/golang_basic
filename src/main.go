package main

import (
	"fmt"
	"./mypackage"
)

func main() {
	var car mypackage.CarPublic
	car.Brand = "Ferrari"
	fmt.Println(car)
	mypackage.PrintCar(car)
}