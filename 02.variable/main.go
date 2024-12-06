package main

import "fmt"

// jwttoken := "sdfdsf" // this will not work outside of function
var jwttoken = "sdfdsf" // this will work

const Pi = 3.1416 // make sure to use capital letter for public constant

func main() {
	var username string = "taohid"
	fmt.Println("Username: ", username)
	fmt.Printf("Variable Type: %T\n",username)

	var isLive bool = true
	fmt.Println("Is Live: ", isLive)
	fmt.Printf("Variable Type: %T\n",isLive)

	var smallInt uint8 = 255
	fmt.Println("Small Int: ", smallInt)
	fmt.Printf("Variable Type: %T\n",smallInt)

	var smalllFloat float32 = 255.234433322564363336
	fmt.Println("Small Float: ", smalllFloat)
	fmt.Printf("Variable Type: %T\n",smalllFloat)

	var malllFloat float64 = 255.234433322564363336
	fmt.Println("Small Float: ", malllFloat)
	fmt.Printf("Variable Type: %T\n",malllFloat)

	// Default Value and some alias

	var name string
	fmt.Println("Name: ", name)
	fmt.Printf("Variable Type: %T\n",name)

	var defaultBool bool
	fmt.Println("Default Bool: ", defaultBool)
	fmt.Printf("Variable Type: %T\n",defaultBool)

	var defaultInt int
	fmt.Println("Default Int: ", defaultInt)
	fmt.Printf("Variable Type: %T\n",defaultInt)

	// imlicit type

	var website = "taohidul.com"
	fmt.Println("Website: ", website)
	fmt.Printf("Variable Type: %T\n",website)

	// no var style

	age := 20
	fmt.Println("Age: ", age)
	fmt.Printf("Variable Type: %T\n",age)

	num := 20.20
	fmt.Println("Num: ", num)
	fmt.Printf("Variable Type: %T\n",num)

	fmt.Println("JWT Token: ", jwttoken)
	fmt.Printf("Variable Type: %T\n",jwttoken)

	fmt.Println("Pi: ", Pi)
	fmt.Printf("Variable Type: %T\n",Pi)

}