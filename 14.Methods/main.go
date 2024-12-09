package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods in Golang")

	avik := User{Name: "Avik",Email: "avik@a.com",Status: true,Age: 25}
	fmt.Printf("User: %+v\n", avik)
	avik.GetStatus()
	avik.NewMail()
	fmt.Printf("User: %+v\n", avik)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)

}

func (u User) NewMail() {
	u.Email = "new" + u.Email
	fmt.Println("New email: ", u.Email)
}
