package main

import (
	"fmt"
)

// course starting with small c means we are not using it outside
// this package and it is only for internal use and if we want to use it outside

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func Json() {
	fmt.Println("Json data in go")
}
