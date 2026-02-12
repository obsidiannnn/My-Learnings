package main

import (
	"fmt"
)

func loops() {
	fmt.Println("loops in GoLang")

	// we will be looping over slice as arrays are not dynamic in golang
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// for loop:

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// for with range

	for i := range days {
		fmt.Println(days[i])
	}

	// loop similar to for each:

	for index, day := range days {
		fmt.Printf("Index is %v and and value is %v \n", index, day)
	}

	value := 1
	for value < 10 {
		if value == 2 {
			goto jump
			// goto is statement by which we can switch to another place
			// like if we define lco somewhere in this and when value hits 2
			// goto will take us to lco and execute that code
		}
		if value == 4 {
			break
		}

		fmt.Printf("Value is %v \n", value)
		value++
	}
	jump: 
	fmt.Println("We jumped to lco when iteration became 2")
	 // after jumping by goto statement our loops breaks automatically
	 


}
