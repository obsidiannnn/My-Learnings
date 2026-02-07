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

}

func moreaboutvar() {
	var name string = "Aditya"

	fmt.Printf("Variable is of type: %T \n", name)

	// fmt.Printf is used to print formatted output
	// \n is neccesary is Printf as it does not go to next line automatically

	// there are various of these predefined variables that are use in formatted output
	// %s is for string

	fmt.Printf("Name: %s \n", name)

	// %d is for interger
	var number int = 34
	fmt.Printf("Number: %d \n", number)

	// %f is for floating-points
	// it by default show upto 6 decimal places

	var price float64 = 99.5
	fmt.Printf("Price %f \n", price)

	//%c is used for writing characters in output

	fmt.Printf("Grade: %c \n", 'A')
	// in this parameter is wrote in single quotes

	// %t is used for bool values.
	var isloggedin bool = true
	fmt.Printf("Login: %t \n", isloggedin)

	// %v is used for values that are given is parameters
	// Safest when you’re not sure about the type.

	fmt.Printf("Anything: %v \n", 42)
	fmt.Printf("Anything: %v \n", "Hi")
	fmt.Printf("Anything: %v \n", true)

	// %T is used to identify type of varibale

	fruit := "mango"
	fmt.Printf("Variable type: %T \n", fruit)
	

}
