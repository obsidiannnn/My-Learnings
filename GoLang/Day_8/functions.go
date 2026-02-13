package main

import "fmt"

func main() {

	fmt.Println("Functions in GO")

	result := adder(24324, 56336)
	fmt.Println(result)

	proresult := proadder(2, 3, 4, 5, 6, 345, 47, 346, 547, 3254, 6, 8, 546, 46)
	fmt.Printf("Pro result is %v \n", proresult)

	Aditya := User{
		"Aditya",
		"adi@email.com",
		true,
		19,
	}
	Aditya.GetStatus()
	Aditya.Newmail()

	def()

	file()
	readfile("./newfile.txt")
}

// taking parameters in a func
// while taking parameters we set type of every value and after that
// we define the type of value that we will return
// just like this a simple int after the bracket

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// when we are not certain how many values will comes
// then we use ... to take all values under common type

//usign ... give us a slice of all values

func proadder(values ...int) int {
	total := 0

	// for takes two values index and value at that index
	// since we not dealing with index in this for
	// that's why we used _

	for _, val := range values {
		total += val
	}
	return total

}
