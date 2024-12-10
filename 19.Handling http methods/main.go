package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Performing a GET request")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	// PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:3000/get" // This is a local server running on port 3000 using express server

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	// content, _ := io.ReadAll(response.Body)
	// fmt.Println("Response Body:", string(content))

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte Count:", byteCount)
	fmt.Println("Response Body:", responseString.String())
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:3000/post" // This is a local server running on port 3000 using express server

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"name": "John Doe",
			"age": 25,
			"isAlive": true
		}
	`)
	fmt.Println(requestBody)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var responseString strings.Builder
	content,_ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte Count:", byteCount)
	fmt.Println("Response Body:", responseString.String())
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:3000/postForm" // This is a local server running on port 3000 using express server

	// form data

	data := url.Values{}
	data.Add("name", "John Doe")
	data.Add("age", "25")
	data.Add("isAlive", "true")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content,_ := io.ReadAll(response.Body)
	fmt.Println("Response Body:", string(content))
}