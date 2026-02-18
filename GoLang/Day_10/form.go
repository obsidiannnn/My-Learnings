package main

import (
	"fmt"
)

func formrequest() {
	fmt.Println("Form request in go ")

	const url = "http://localhost:8000/postform"

	// fake form data

	data := url.Values{}

	// url.Values is a type that represents the form data
	// as key-value pairs.

	data.Add("firstname", "Aditya")

	// data.Add is used to add key-value pairs to the form data.

	data.Add("lastname", "Chauhan")
	data.Add("email", "adi@gmail.com")

}
