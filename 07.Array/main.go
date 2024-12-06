package main

import "fmt"

func main() {
	fmt.Println("Array in Golang")

	// Array Declaration
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Banana"
	fmt.Println("Fruit List is: ", fruitList)

	fmt.Println("the length of the fruit list is: ", len(fruitList))

	// Array Declaration and Initialization
	var vegList = [3]string{"Potato", "Beans", "Onion"}
	fmt.Println("Vegetable List is: ", vegList)
	
} 