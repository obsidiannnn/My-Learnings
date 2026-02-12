package main

import (
	"fmt"
)

func main(){
	fmt.Println("Functions in GO")

	result := adder(24324,56336)
	fmt.Println(result)
	
}

// taking parameters in a func
// while taking parameters we set type of every value and after that
// we define the type of value that we will return
// just like this a simple int after the bracket

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}