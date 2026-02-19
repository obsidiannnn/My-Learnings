package main

import (
	"encoding/json"
	"fmt"
)

// course starting with small c means we are not using it outside
// this package and it is only for internal use and if we want to use it outside

// structs are often used to represent data in jsson format
// and we can use json tags to specify the json key for each field in the struct

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

func Encodingjson() {
	// mycourse is a slice of course struct and it is empty for now
	// and we will add some data to it later.
	mycourse := []course{
		{"ReactJS", 0, "online", "abc123", []string{"web-dev", "frontend"}},
		{"GoLang", 0, "online", "def456", []string{"web-dev", "backend"}},
		{"Python", 0, "online", "ghi789", nil},
	}

	// package this data as json data and send it to client

	// json data is a format that is used to send data over the internet

	//finaljson, err := json.Marshal(mycourse)

	// while using json.Marshal it will return data in clustered format
	// and to make it more readable we can use json.MarshalIndent

	finaljson, err := json.MarshalIndent(mycourse, "", "\t")

	// here "" means we don't want any prefix and "\t" means we
	// want to use tab for indentation and it will make data more readabl.

	if err != nil {
		panic(err)
	}

	// json.Marshal is used to convert the data into json format
	// and it returns the json data in byte format and error if any.

	fmt.Printf("final json is %s\n", string(finaljson))
}
