package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("GET request in GO")

	getrequest()

	post()
}

func getrequest() {
	const newurl = "http://localhost:8000/get"

	res, err := http.Get(newurl)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// once again api request never close on their own it is imp to
	// close them at last.

	fmt.Println("Status code: ", res.Status)
	fmt.Println("Content length is: ", res.ContentLength)

	content, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Content of response is: ", string(content))

	// content is in byte format to we use string to make it readable



	
}
