package main

import (
	"fmt"
)

func main() {
	maps()
}

func maps() {
	lang := make(map[string]string)

	// first string is for dtype of keys and second is for values

	lang["JS"] = "Javascript"
	lang["RB"] = "Ruby"
	lang["PY"] = "Python"

	// maps in go are only for data accessing
	// don't rely on it's order because it doesn't have it

	fmt.Println(lang)

	// accessing via key

	fmt.Println("JS stands for: ", lang["JS"])

	// delete keyword can be used to remove any value

	delete(lang, "RB")
	fmt.Println(lang)

	// loops in maps:

	// method to write for loop for map

	for key, value := range lang {
		fmt.Printf("For Key %v, value is %v", key, value)
	}

}
