package main

import "fmt"

func variables() {
	//var name type
	var name string
	name = "Aditya"

	var year int = 2026

	// we directly assign value to a variable
	// after defining it's type

	fmt.Println("Name: ", name)
	fmt.Println("Year: ", year)

	fmt.Printf("Variable is of type: %T", name)

}
