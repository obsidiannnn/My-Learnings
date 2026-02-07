package main

import (
	"fmt"
	"sort"
)

func slicing() {
	var item = []string{"Iphone", "macbook", "PS5", "fan", "Bat"}

	// adding items by slicing

	item = append(item, "Cap")

	fmt.Println(item[0:])

	fmt.Println(item[1:9])

	// while accessing like this since there are not 9 values
	// some space will be generated after last element
	// instead of throwing an error.

	// slicing in go lang is similar to python
	// but in python it will throw an error if we try to access
	// more than available elements

	fmt.Println(item[2:4])
	fmt.Println(item[2:2]) // will return empty list

	highscores := make([]int, 4)

	// make makes a slice of given size which we can fill later
	// it doesn't return any pointer it returns the slice itself
	// same goes for variable made with make function

	highscores[0] = 234
	highscores[1] = 456
	highscores[2] = 678
	highscores[3] = 890

	highscores = append(highscores, 1000, 2000, 3000)

	// since we allocate memory of 4 still we are able to add more elements
	// beacause append will reallocate the memory and copy the old elements

	fmt.Println(highscores)

	// there are various packages avaliable for slices
	// for different operations like sort

	sort.Ints(highscores)
	// sort in asecnding order

	fmt.Println(highscores)

	// we can also get boolean values with this like we can ask is sorted or not?

	fmt.Println(sort.IntsAreSorted(highscores))

	// will return true cz it is sorted

	// removing value from slice based on index

	var index int = 2

	// we can use append to remove values too
	// logic is see we took index as 2 to remove that ele
	// we need ele before index and after index since last ele is excluded first range will be till index
	// and next will be index +1 till last

	highscores = append(highscores[:index], highscores[index+1:]...)
	fmt.Println(highscores)

}
