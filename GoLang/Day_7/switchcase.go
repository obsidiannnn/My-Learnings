package main

import (
	"fmt"
	"math/rand"
	"time"
)

func switchcase() {
	fmt.Println("switch case in golang")

	rand.Seed(time.Now().UnixNano())
	dicenum := rand.Intn(6) + 1

	// for rolling a die we need six digit but with random
	// last digit is not included ie 6+1 or 7

	fmt.Println("Value of Dice is: ", dicenum)

	switch dicenum {
	case 1:
		fmt.Println("Dice Value is one and you can open")
	case 2:
		fmt.Println("Move 2 spots ahead")
	case 3:
		fmt.Println("Move ahead by 3 spots")
		fallthrough

		// fall through means with case 3 occurs run the case
		// just after 3 too that means case will also run
		// note only just after case will run in fallthrough
		
	case 4:
		fmt.Println("Move ahead by 4 spots")
	case 5:
		fmt.Println("Move ahead by 5 spots")
	case 6:
		fmt.Println("Move ahead by 6 spots and roll the die again")
	default :
		fmt.Println("What was that")

	// for edge cases if none of the above case matches 
	// default state runs

	}

}
