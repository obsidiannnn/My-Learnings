package main

import (
	"fmt"
	"strings"
)

func post() {
	fmt.Println("POST request in golang")

	const urll = "http://localhost:8000/post"

	// fake json payload , payload is known as sending some json data

	requestBody := strings.NewReader(`
		{
			"coursename" : "GoLang",
			"price" : 0,
			"platform" : "online"
		}
	`)

	res, err = http.POST(urll, "application/json", requestBody)
}
