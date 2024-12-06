package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Handling in Go")
	presentTime := time.Now()
	fmt.Println("The current time is:", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2024, time.December, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("The course was created on:", createdDate.Format("01-02-2006 Monday"))
}
