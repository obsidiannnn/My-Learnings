package main

import (
	"fmt"
)

// if u want to create a global/public variable in go
// Variable name must start with capital letter
// and no inferred var declaration is allowed .

var Public bool = true

func shortdeclaration() {
	var city string
	city = "London"

	var name = "Aditya" // inferred to string

	// i.e. var datatype will be declared automatically
	// Correct way for shortdeclaration:

	currentmoney := 5000

	currentmoney += 1000

	fmt.Println(city, name, currentmoney)
	fmt.Println(Public)

}
