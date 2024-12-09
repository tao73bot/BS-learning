package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	// For loop
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i:=0; i<len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i:= range days {
	// 	fmt.Println(days[i]) // i returns index
	// }

	for index, day := range days {
		fmt.Printf("Index: %v, Day: %v\n", index, day)
	}

	x := 1

	for x < 10 {
		if x == 5 {
			x++
			continue
		}
		if x == 8 {
			goto lco
		}
		fmt.Println(x)
		x++
	}

	lco :
	     fmt.Println("Jumping to LCO")
}