package main

import (
	"fmt"
	"io"
	"net/http"
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

	res, err := http.Post(urll, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("response after posting data is: ", string(content))


}
