package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to conversion in Go")
	fmt.Println("Please rate this course [1-5]:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	num, err := strconv.ParseInt(strings.TrimSpace(input),0 ,64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Thanks for giving us :", num)
	}
}
