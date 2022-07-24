package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizz world")
	fmt.Println("Enter your Rating between 1 and 5 ")

	scanner := bufio.NewReader(os.Stdin)  //initialize the scanner
	rating, _ := scanner.ReadString('\n') //use scanner to read string and store the value in rating variable

	println("Thankyou for you rating of ", rating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64) //TrimSpace is to remove \n from scanned input
	if err != nil {
		fmt.Println(err)
		//panic(err) is used if you want to terminate your program if err occurs
	} else {
		fmt.Println("Your incremented rating is ", numRating+1)
	}
}
