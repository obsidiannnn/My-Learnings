package main

import (
	"fmt"
)

func Slice() {

	var items = []string{"One", "Two", "Three"}

	// while defining if we give fixed number of elements it's a array
	// if we don't give then it's a slice

	// in slice we can add and subtract as many as values as we like

	fmt.Println(items)
	fmt.Printf("Type of itmes is %T \n", items)

	// Method to add to string at the end:

	items = append(items, "Four", "Five")
	fmt.Println(items)

}
