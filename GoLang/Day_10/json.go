package main

import (
	"fmt"
	"encoding/json"
)

// course starting with small c means we are not using it outside
// this package and it is only for internal use and if we want to use it outside

type course struct {
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string

	// tag will be a slice of string and it will 
	// hold multiple tags for the course
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
	// /

	if err != nil {
		panic(err)
	}

	// json.Marshal is used to convert the data into json format 
	// and it returns the json data in byte format and error if any.

	fmt.Printf("final json is %s\n", string(finaljson))
}
