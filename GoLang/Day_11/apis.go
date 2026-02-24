package main

import (
	"fmt"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	Coursename  string  `json:"coursename"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"`

	// we are defining author as custom type were we defined it's type elsewhere
	// like struct and used here now while taking author values we certainly
	// don't want to reference of value but actual values , that's why we use
	// pointers.

}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB

// it will be a slice of type course

// if courseid is not avaliable we will not move further.

var courses []Course



func main() {
	fmt.Println("Building api in go")
}
