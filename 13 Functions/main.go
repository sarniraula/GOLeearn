package main

import "fmt"

func main() {
	fmt.Println("Functions in go Lang")
	greeter()

	result := adder(3, 5)

	proRes, _ := proAdder(3, 4, 5, 8, 56)

	fmt.Println("Result is: ", result)
	fmt.Println("Pro adder result is: ", proRes)

	// func greeeterTwo() {
	// 	fmt.PrintLn("Second greeting")  NOT ALLOWED FUNCTION WITHIN ANTOHER FUNCTION
	// }
}

func greeter() {
	fmt.Println("Namastee")
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hello ProAdder Function"
}
