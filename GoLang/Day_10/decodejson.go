package main

import (
	"fmt"
)

type course struct {
	Name string `json:"Coursename"` // it will replace field name with coursename
	// all over the data and it will be used as key in json data instead of Name

	Price    int
	Platform string `json:"website"`
	Password string `json:"-"` // after putting dash it will not be
	// included in json data and it will be ignored while encoding
	Tags []string `json:"tags,omitempty"` // after putting omitempty
	// if the value of tags is empty then it will not be included in json data

	// and it will return null as output.
}

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

	var mycourse course

	// json data that is coming from has to be stored in some
	// variable and that variable should be of type course struct

	// we can check wheather our data is in correct format or not
	// by using json.Valid function

}
