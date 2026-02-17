package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("GET request in GO")

	getrequest()
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
	fmt.Println("Content length is: ",res.ContentLength)

}
