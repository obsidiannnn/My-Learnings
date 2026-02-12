package main

import (
	"fmt"
)

func loops(){
	fmt.Println("loops in GoLang")

	// we will be looping over slice as arrays are not dynamic in golang
	days := []string{"Sunday","Tuesday","Wednesday","Friday","Saturday"}
	fmt.Println(days)

	// for loop:

	for d:= 0; d<len(days); d++{
		fmt.Println(days[d])
	}

	// for with range

	for i := range days{
		fmt.Println(days[i])
	}
}