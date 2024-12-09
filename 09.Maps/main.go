package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maps in Golang!")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "Golang"
	languages["TS"] = "TypeScript"
	languages["JAVA"] = "Java"
	languages["CS"] = "C#"
	languages["CPP"] = "C++"
	languages["R"] = "R"


	fmt.Println("List of languages:", languages)
	fmt.Println("JS:", languages["JS"])

	// Delete a key
	delete(languages, "RB")
	fmt.Println("List of languages after deleting Ruby:", languages)

	// loop through map
	for key, value := range languages {
		fmt.Println("for", key, "we have", value)
	}
}
