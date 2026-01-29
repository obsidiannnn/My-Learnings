package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func conv() {
	fmt.Println("Welcome to converison")
	fmt.Println("Enter rating for food")

	reader := bufio.NewReader(os.Stdin)

	// reader is just made to read what user type
	// what reader will read we will define here

	input, _ := reader.ReadString('\n')
	// it will read input as string

	fmt.Println("Thanks for rating, ", input)

	// now to convert this into int we use package strconv
	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// in this we give two value after telling to which we have to convert
	// like input we have to conver and to float64

	// in this we have to face error as well

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added one to your rating,", numrating+1)

		//strconv.ParseFloat: parsing "55\n": invalid syntax
		// this error is coming in terminal which means there is
		// extra character that is space that is \n in input
		// that can be convert
		// so that we use another package strings in which we
		// use strings.trimspace(input) to remove that
	}
}
