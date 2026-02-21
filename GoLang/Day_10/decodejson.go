package main

import (
	"fmt"
)

func DecodeJson() {
	fmt.Println("Decodind json data")

	// when data comes from net it is in byte format to amke it readble
	// we use strings
	// but sometimes we need data as it is in json format that's why we use
	// []byte

	jsondata := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"Price" : 299,
			"Website": "HelloWorld.com",
			"tags": ["web-dev","js"]
		}
		{
			"coursename": "GoLang Bootcamp",
			"Price" : 0,
			"Website": "HelloWorld.com",
			"tags": ["web-dev","go"]
		}
	
	`)
}
