package main

import (
	"fmt"
)

func main(){
	ifelse()
}

func ifelse(){

	var result string
	logincount :=  25

	if logincount >10{
		result  = "Regular User"
	} else if logincount < 10 {
		result = "New User"
	}
	fmt.Println(result)


}