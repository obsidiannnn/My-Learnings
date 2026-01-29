package main

import (
	"bufio" // bufio and os are two packages used to take
	"fmt"
	"os" // user input
)

func input() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for food") // message for user for input

	// Here we are using walrus operator for reader as
	// we don't know what type of data will come
	// so it's data type will be automatically decided

	// we are creating a Newreader which will take user input
	// stdin is standard input output of package os

	// now we want to store reader ans in varibale, here we use
	// comma ok / err ok syntax
	// in go there is no try catch so if any error comes up we have to deal with it in from of varibales and
	// boolean value

	input, _ := reader.ReadString('\n')
	/// we are taking input and _ as error
	// in go if u not certain about any varibale u can just take
	// _ in place of that varibale
	// if I had took error as varibale and didn't use it
	// it has given me an error
	// \n means we want input till \n ie till enter

	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of this rating %T \n", input)

	// currently if I am typing anything as user it is stored as string

}
