package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")
    var a *int
	fmt.Println("Value of a: ", a)
	fmt.Printf("Type of a: %T\n", a)

	num := 23
	var ptr = &num // & operator is used to get the address of the variable
	fmt.Println("Value of num: ", num)
	fmt.Println("Address of num: ", &num)
	fmt.Println("Value of ptr: ", ptr)
	fmt.Println("Value of *ptr: ", *ptr) // * operator is used to get the value stored at the address

	* ptr = *ptr * 2
	fmt.Println("Value of num: ", num)
	fmt.Println("Value of *ptr: ", *ptr)
	
}