package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch in go lang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is one, and you can open")
	case 2:
		fmt.Println("Dice value is two")
	case 3:
		fmt.Println("Dice value is three")
	case 4:
		fmt.Println("Dice value is four")
	case 5:
		fmt.Println("Dice value is five")
	case 6:
		fmt.Println("Dice value is six and you can roll again")
	default:
		fmt.Println("What was that!!")
	}
}
