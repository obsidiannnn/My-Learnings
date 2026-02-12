package main

import (
	"fmt"
)

// func execution happens line by line
// but if we use defer keyword that line goes to last line of the func
// ie whole function will be executed first then the defer lines will be executed

// defer works as last in first out
// ie if there are many defer lines then the lines written at the last
// will be executed first

func def(){
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	mydefer()

	// in this number from 1 to 4 will be stored in stack
	// as it is a defer statement then acc to last in first out
	// numbers will be printed as 4,3,2,1 then rest of the code accordingly.
	
}

func mydefer(){
	for i := 0; i<5; i++ {
		defer fmt.Println(i)
	}
}

