package main

import "fmt"

func main() {
	fmt.Println("Welcome to the study on pointers")

	myNumber := 25

	var ptr = &myNumber

	fmt.Println("Address of my number is, ", ptr)
	fmt.Println("Actual value of my number is, ", *ptr)

	*ptr = *ptr * 2

	fmt.Println("My number doubled is, ", myNumber) //Actual value of myNumber variable doubled using ptr
}
