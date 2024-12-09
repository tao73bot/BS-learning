package main

import "fmt"

func main() {
	
	defer fmt.Println("This is a deferred statement")
	defer fmt.Println("This is another deferred statement") // LIFO
	fmt.Println("Welcome to Defer in Golang")
	defer fmt.Println("This is the last deferred statement")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}