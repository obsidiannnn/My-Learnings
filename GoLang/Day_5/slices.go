package main

import (
	"fmt"
)

func Slice() {

	var items = []string{"One", "Two", "Three"}

	// while defining if we give fixed number of elements it's a array
	// if we don't give then it's a slice

	fmt.Println(items)
	fmt.Printf("Type of itmes is %T \n", items)
}
