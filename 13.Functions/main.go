package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")
	avik()
	fahad()
	// we can not define a function inside a function
	// func fahad() {
	// 	fmt.Println("Fahad the Go programmer")
	// }
	// fahad()
	result := add(10, 20)
	fmt.Println("Result is: ", result)
	proResult := proAdd(10, 20, 30, 40, 50)
	fmt.Println("Pro result is: ", proResult)
}

func avik() {
	fmt.Println("Tre pani ditam na!")
}

func fahad() {
	fmt.Println("Fahad the Go programmer")
}

func add(x int, y int) int {
	return x + y
}

func proAdd(value ...int) int {
	total := 0
	for _, val := range value {
		total += val
	}
	return total
}