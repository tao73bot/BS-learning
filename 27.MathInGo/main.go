package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to math in Go!")

	var num1 int = 2
	var num2 float64 = 4.5
	fmt.Println("Addition: ", num1+int(num2))

	// random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random number using crypto/rand
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}
