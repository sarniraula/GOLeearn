package main

import "fmt"

func main() {
	fmt.Println("Welcome to the study on Slices")

	var fruitlist = []string{"Apple", "Banana", "Peach"}
	fmt.Printf("The type of fruitlist is %T \n", fruitlist)
	fmt.Println(fruitlist)
	fmt.Println("THe length of slice fruitlist is: ", len(fruitlist))

	fruitlist = append(fruitlist, "mango", "watermelon") //To add items to the slices
	fmt.Println(fruitlist)
	fmt.Println("THe updated length of slice fruitlist is: ", len(fruitlist))

	var courses = []string{"Reactjs", "Python", "Swift", "C", "Ruby"}
	fmt.Println(courses)
	var index int = 2

	//Removing the index element"swift"" from the slice
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
