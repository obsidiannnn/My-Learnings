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

	http.Get(newurl)
}
