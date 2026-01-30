package main

import (
	"fmt"
	"time"
)

func main() {
	Time()
	pointers()

}

func Time() {
	fmt.Println("Time in Go")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// This is standard input for format no matter what day, time and date it is u have to
	// pass these only after that go will return today's date, time and day

	createdate := time.Date(2020, time.March, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdate)
	fmt.Println(createdate.Format("01-02-2006 Monday"))
}
