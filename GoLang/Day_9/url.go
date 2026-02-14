package main

import (
	"fmt"
	"net/url"
)

// creating A demo url
const myurl string = "https://demourl:3000/learn?coursename=golang&paymentid=ghbj456ghb"

func urll(){
	fmt.Println("Handling urls in golang")
	fmt.Println(myurl)

	// parsing the url - means extracting different parts from url
	
	result, err :=  url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme is: ", result.Scheme)
	fmt.Println("Host is: ", result.Host)
	fmt.Println("Path is: ", result.Path)
	fmt.Println("Port is: ", result.Port())
	fmt.Println("Query is: ", result.Query())
	fmt.Println("Raw Query is: ", result.RawQuery)

}