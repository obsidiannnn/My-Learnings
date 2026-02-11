package main

import (
	"fmt"
)

func Struct() {
	fmt.Println("Structs in GoLang")
	// there is no inheritance in golang
	// or super or parent

	// type is keyword, User is name , and we tell this is a
	// struct by writing struct

}

type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}

	// User U is capital , same way all fields start from capital letters
	// as these are alternative of classes so they need to be export
	// capital alphabet in starting means it can be used anywhere
	// is it starts with samll letter means it can be user here only

	// a struct should be  defined outside a func otherwise it
	// scope will be only that func and not globally
	

