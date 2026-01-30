package main

import "fmt"

// When passing a varibale sometimes copy of these variables are passed on
// to aviod such things pointers are used
// Pointer is reference direct to the memory address
// which makes sure that the actual value is passed on

func pointers() {
	var ptr *int

	// * is used to create a pointer
	// by default the value of pointer varibale with no value is nil

	fmt.Println(ptr)

	num := 23

	var pointr = &num

	// here we are making a pre defined varibale as pointer
	// so we used & instead of *

	fmt.Println("Value of pointer is ", pointr)
	fmt.Println("Actual Value is ", *pointr)

	// while calling simply pointer it will give the  address
	// to which it is refereing to
	// to get it actual value use *pointer or varible name

	*pointr = *pointr + 2

	// pointer allow these type of operations
	fmt.Println("New Value: ", *pointr)

	// while doing operations like this, operations are perfomred on actual
	// values and not on the copy of the values

}
