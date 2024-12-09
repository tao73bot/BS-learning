package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Go")

	// Create a file
	content := "This is a file created from Go"
	file, err := os.Create("./myFile.txt")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length of the file is: ", length)
	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(filename string) {
	dataByte,err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data in the file is: ", string(dataByte))
}
