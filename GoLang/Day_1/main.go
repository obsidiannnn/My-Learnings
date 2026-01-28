package main

//gettings starting with go

import (
	"fmt"
)

// main() is the entry point
// run a go file -> go looks for package main and func main

func main() {
	fmt.Println("Hello Go")

	Math()

	// We can directly use other go files in main file if
	// they are in same folder and are of same package

	variables()

	moreaboutvar()
}

// Inside a go folder there will be one main.go file
// that will run the code , we will code in other files
// to but have to call those fnc inside the main files's
// func main to run them.
