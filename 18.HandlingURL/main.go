package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://loc.dev:3000/learn?courename=reactjs&payment=ghbj33431"

func main() {
	fmt.Println("Handling URL")

	fmt.Println("URL:", myurl)

	// Parse the URL
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("RawQuery:", result.RawQuery)
	fmt.Println("Query:", result.Query())
	fmt.Println("Port:", result.Port())
	fmt.Println("Fragment:", result.Fragment)

	qparams := result.Query()

	fmt.Println("Course Name:", qparams["courename"])
	fmt.Println("Payment:", qparams["payment"])

	for _,value := range qparams {
		fmt.Println("Value:", value)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "loc.dev:3000",
		Path: "/learn",
		RawQuery: "course=reactjs&payment=ghbj33431",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Another URL:", anotherUrl)
}