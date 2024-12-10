package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"course_name"`
	Price    float64  `json:"price"`
	Platform string   `json:"platform"`
	password string   `json:"-"` // this field will be ignored
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("This is a JSON example in Go")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	taohidCourse := []course{
		{"React Native", 7.99, "Udemy", "122", []string{"Mobile", "React"}},
		{"Go", 57.99, "Udemy", "233", []string{"Backend", "Go"}},
		{"Python", 51.99, "Udemy", "344", nil},
		{"Docker", 37.99, "Udemy", "455", []string{"DevOps", "Docker"}},
	}

	// package this data into a JSON
	finalJson, err := json.MarshalIndent(taohidCourse, "", "\t")
	if err != nil {
		fmt.Println("Error in marshalling JSON")
		return
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"course_name": "React Native",
		"price": 7.99,
		"platform": "Udemy",
		"tags": ["Mobile", "React"]
	}
	`)

	var courseData course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON is valid")
		err := json.Unmarshal(jsonDataFromWeb, &courseData)
		if err != nil {
			fmt.Println("Error in unmarshalling JSON")
		}
		fmt.Printf("%+v\n", courseData)
	}else {
		fmt.Println("JSON is not valid")
	}

	// some cases where you just want to add data to key value pair

	var myOnlineData map[string]interface{}
	err := json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	if err != nil {
		fmt.Println("Error in unmarshalling JSON")
	}
	fmt.Printf("%#v\n", myOnlineData)
	for key, value := range myOnlineData {
		fmt.Printf("Key: %v, Value: %v, Type of value: %T\n", key, value,value)
	}
}
