package main

import (
	"fmt"
)

// since methods are way to deal with classes and in go there are no
// classes instead there are structs

type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}

// way to define a method in go

// u is a alias name for user
func (u User) GetStatus() {
	fmt.Printf("Is User active: %v \n", u.Status)
}

func (u User) Newmail() {
	fmt.Printf("Previos mail : %v \n", u.Email)
	u.Email = "hello@mail.com"
	fmt.Printf("New mail after manupilating : %v \n",u.Email)

	// changing email like this doesn't change the original email
	// only reference is passed that's why we use pointers

}