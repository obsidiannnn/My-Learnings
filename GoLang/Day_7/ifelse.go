package main

import (
	"fmt"
)

func main(){
	ifelse()
	switchcase()
}

func ifelse(){

	var result string
	logincount :=  25

	if logincount >10{
		result  = "Regular User"
	} else if logincount < 10 {
		result = "New User"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Not even")
	}

	// initializing and checking for the condition at the same time

	if num := 3; num <10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Greater than 10")
	}
}