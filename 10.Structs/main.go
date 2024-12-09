package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")

	// no inheritance in Go
	// no super or parent keyword in Go
	// no subclass or child keyword in Go

	avik := User{Name: "Avik", Email: "avik@a.com", Status: true, Age: 25}
	fmt.Println("Avik: ", avik)
	fmt.Printf("Avik details: %+v\n", avik)

	// see single field
	fmt.Printf("Name is %v and Email is %v\n", avik.Name, avik.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
