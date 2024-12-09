package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://go.dev/tour/list"

func main() {
	fmt.Println("Web requests")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	datatypes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data is of type: %T\n", datatypes)
	fmt.Println(string(datatypes))
}