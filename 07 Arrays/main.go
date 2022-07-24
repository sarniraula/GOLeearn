package main

import "fmt"

func main() {
	fmt.Println("welcome to arrays in Go Lang")

	var fruitList [4]string //in arrays no of items must be provided compulsary
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("Length of fruit list arrray; ", len(fruitList)) //No matter what you do the length is 4

	var veglist = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("vegetable list is: ", veglist)
	fmt.Println("Length of veglist: ", len(veglist))
}
