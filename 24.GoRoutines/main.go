package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointer to the WaitGroup struct
var mut sync.Mutex    // pointer to the Mutex struct

func main() {
	fmt.Println("Welcome to GoRoutines")
	// go greeter("Hello")
	// greeter("World")
	websiteList := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.golang.org",
		"https://www.medium.com",
		"https://www.microsoft.com",
		"https://www.linkedin.com",
	}
	for _, v := range websiteList {
		go getStatusCode(v)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("Signals:", signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3*time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Println("Status Code for", endpoint, "is", res.StatusCode)
	}
}
