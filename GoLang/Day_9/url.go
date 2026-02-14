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
	// output https

	fmt.Println("Host is: ", result.Host)
	// output demourl:3000

	fmt.Println("Path is: ", result.Path)
	// output /learn
	fmt.Println("Port is: ", result.Port())
	// output 3000, port is not a query it is a method to get port from url

	fmt.Println("Query is: ", result.Query())
	// output map[coursename:[golang] paymentid:[ghbj456ghb]]
	
	fmt.Println("Raw Query is: ", result.RawQuery)
	// output coursename=golang&paymentid=ghbj456ghb

	// to get value of a particular query parameter
}