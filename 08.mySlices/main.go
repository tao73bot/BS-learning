package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	// Declaring a slice
	var fruitList = []string{}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Apple", "Tomato", "Banana")
	fmt.Println("Fruit List is: ", fruitList)
	fruitList = append(fruitList, "Mango")
	fmt.Println("Fruit List is: ", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("Fruit List is: ", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 234
	highScores[2] = 345
	highScores[3] = 456

	// highScores[4] = 555 // This will give an error
	highScores = append(highScores, 555) // This will work
	// 4 is the inicial length of the slice and using append we can add more elements to the slice and the length will increase

	fmt.Println("High Scores: ", highScores)
	fmt.Printf("Length of highScores: %d\n", len(highScores))
	fmt.Printf("Capacity of highScores: %d\n", cap(highScores))
	highScores = append(highScores, 666, 777, 888, 999)
	fmt.Println("High Scores: ", highScores)
	fmt.Printf("Length of highScores: %d\n", len(highScores))
	fmt.Printf("Capacity of highScores: %d\n", cap(highScores))


	sort.Ints(highScores)
	fmt.Println("Sorted High Scores: ", highScores)
	fmt.Printf("Are high scores sorted? %v\n", sort.IntsAreSorted(highScores))

	// removing an element from a slice based on index
	var courses = []string{"react", "angular", "vue", "svelte","nodejs", "flutter"}
	fmt.Println("Courses: ", courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses after removing: ", courses)
}
