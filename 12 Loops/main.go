package main

import "fmt"

func main() {
	fmt.Println("Loops in Go lang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	for i := 0; i < len(days); i++ { //Type 1
		fmt.Println(days[i])
	}

	for d := range days { //Type 2
		fmt.Println(days[d])
	}

	for d, day := range days { //Type 3 (for each) can also use comma ok syntax as  "for _,day := range .."
		fmt.Println(d, day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 5 {
			goto jumpHere
		}
		fmt.Println("Value is ", rogueValue)
		rogueValue++
	}

jumpHere:
	fmt.Println("Jumped here")
}
