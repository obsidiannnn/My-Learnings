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
}
