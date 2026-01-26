package main

import (
	"fmt"
)

func shortdeclaration() {
	var city string
	city = "London"

	var name = "Aditya" // inferred to string

	// i.e. var datatype will be declared automatically
	// Correct way for shortdeclaration:

	currentmoney := 5000

	currentmoney += 1000

	fmt.Println(city, name, currentmoney)

}
