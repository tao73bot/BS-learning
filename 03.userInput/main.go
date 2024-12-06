package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the playground!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	// comma ok idiom
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, " + name)
	fmt.Printf("Type of input, %T\n", name)

	age, _ := reader.ReadString('\n')
	fmt.Println("Age: ", age)
	fmt.Printf("Type of input, %T\n", age)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Input: ", input)
		fmt.Printf("Type of input, %T\n", input)
	}
}
