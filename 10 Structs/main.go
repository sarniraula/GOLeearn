package main

import "fmt"

type User struct { //User with capital U
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in Go Lang")
	//No inheritence in go lang; No super or parent or child

	saroj := User{"Saroj", "sar.niraula@gmail.com", true, 29}

	fmt.Println(saroj) //Output => {Saroj sar.niraula@gmail.com true 29}

	fmt.Printf("Saroj's details are: %+v\n", saroj) //Output => Saroj's details are: {Name:Saroj Email:sar.niraula@gmail.com Status:true Age:29}

	fmt.Printf("Name is %v and Email is %v.", saroj.Name, saroj.Email) //Output => Name is Saroj and Email is sar.niraula@gmail.com.

}
