package main

import (
	"fmt"
	"io"
	"os"
)

func file(){
	fmt.Println("Handling files in go")
	content := "This needs to go in file"

	// os is a module used to work with operating system
	// we can use this to create a file.

	file, err := os.Create("./newfile.txt")

	// this is comma ok syntax means files can we created or
	// there might be some error so we will write code to handle err too.

		if err != nil {
			panic(err)
		}
		// we check if err has something or not if it has something
		// we use panic, panic basically shut downs the execution of
		// program and return the error. 

		length, err := io.WriteString(file,content)

		// io package is used to write in file 
		// so we gave two parameters first in which file we have to write
		// second what we have to write.

		// we can check that if something is written in file or not
		// if it is then it will return length of that string otherwise error.

		if err != nil{
			panic(err)
		}
		fmt.Println("Length is: ", length)
		defer file.Close() //! to close the file

		// we used defer cz we might need some code to write after this too
		// to obv file should be closed at the end
}

// creating a func to read file

func readfile(){
	
}