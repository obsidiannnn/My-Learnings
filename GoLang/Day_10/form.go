package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func formrequest() {
	fmt.Println("Form request in go ")

	const li = "http://localhost:8000/postform"

	// fake form data

	data := url.Values{}

	// url.Values is a type that represents the form data
	// as key-value pairs.

	data.Add("firstname", "Aditya")

	// data.Add is used to add key-value pairs to the form data.

	data.Add("lastname", "Chauhan")
	data.Add("email", "adi@gmail.com")

	res, err := http.PostForm(li, data)

	// http.PostForm is used to send a POST request with form data
	// and it takes the url and form data as parameters.

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("response after posting form data is: ", string(content))
}
