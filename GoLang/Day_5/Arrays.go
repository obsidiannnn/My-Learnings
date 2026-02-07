package main

import (
	"fmt"
)

func main() {
	Arr()
	Slice()
	slicing()

}

func Arr() {
	var fruit [4]string

	// in arrays yo have to tell the number of elements and type of
	// array in starting, then add values

	fruit[0] = "Apple"
	fruit[1] = "Mango"
	fruit[2] = "Peach"
	fruit[3] = "Banana"

	fmt.Println(fruit)
	fmt.Println(len(fruit))

	// While calculating the len no matter how many values are there
	// len while return the len that was decided while declaring
	// of array irrespective of the elements in it.

	var list = [3]string{"Hello", "Hi", "Bye"}

	// Another method to declare

	fmt.Println(list)
	fmt.Println(len(list))
}
