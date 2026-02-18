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

	// while using post request it goes in application/json format
	// and we have to specify the content type as well and 
	// then we have to pass the request in request body.

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)

	// io.ReadAll is used to read the response body 
	// and it returns the content in byte format and error if any.

	if err != nil {
		panic(err)
	}

	fmt.Println("response after posting data is: ", string(content))


}
