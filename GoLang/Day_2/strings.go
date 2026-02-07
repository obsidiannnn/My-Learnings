package main

import (
	"fmt"
	"strings" // import string package to use it's
	// pre defiened features.
)

func Strings() {
	firstname := "Aditya"
	lastname := "Chauhan"
	fullname := firstname + "" + lastname

	fmt.Println(fullname)

	fmt.Println(strings.ToUpper(fullname))
}
