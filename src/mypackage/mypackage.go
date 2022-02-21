package mypackage

import (
	"fmt"
)
//CarPublic car con accesso publico
type CarPublic struct{
	Brand string
	Year int
}

func PrintCar(car CarPublic){
	fmt.Println(car)
}