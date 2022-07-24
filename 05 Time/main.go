package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang")
	presentTime := time.Now()
	fmt.Println(presentTime) //Prints time in default format UGLY
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	createdDate := time.Date(2022, time.April, 21, 04, 05, 55, 0, time.UTC) //Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) time.Time
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("Monday 01/02/2006 15:04:05 pm")) //Must use the same date 01-02-2006 Monday 15:04:05 always
}
