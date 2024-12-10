package main

import (
	"fmt"
	"log"
	"net/http"
	"realDb/router"
)

func main() {
	fmt.Println("Working with RealDB")
	r := router.Router()
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server stopped")
}
