package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course
type Course struct {
	CourseId    string  `json:"courseId"`
	Title       string  `json:"title"`
	CoursePrice float64 `json:"coursePrice"`
	Author      *Author `json:"author"`
}

// fake DB
var courses []Course

// middleware, helper function
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.Title == ""
	return c.Title == ""
}

// model for author
type Author struct {
	FullName string `json:"fullname"`
	Username string `json:"username"`
	WebSite  string `json:"website"`
}

func main() {
	fmt.Println("Building RESTful API in Go")
	r := mux.NewRouter()
	// seeding
	courses = append(courses, Course{CourseId: "1", Title: "Course 1", CoursePrice: 9.99, Author: &Author{FullName: "John Doe", Username: "johndoe", WebSite: "https://johndoe.com"}})
	courses = append(courses, Course{CourseId: "2", Title: "Course 2", CoursePrice: 19.99, Author: &Author{FullName: "Jane Doe", Username: "janedoe", WebSite: "https://janedoe.com"}})
	courses = append(courses, Course{CourseId: "3", Title: "Course 3", CoursePrice: 29.99, Author: &Author{FullName: "Jim Doe", Username: "jimdoe", WebSite: "https://jimdoe.com"}})

	// routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":8080", r))
}

// controller
// server home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the home page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// get the course id from the url
	params := mux.Vars(r)

	// loop through the courses to find the course with the id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("No course found with that id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// if the request body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send a course")
		return
	}

	// if send data is like {}
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid course data")
		return
	}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send a valid course")
		return
	}

	// generate an unique course id
	// append the course to the courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// get the id from the url
	params := mux.Vars(r)

	// loop through the courses to find the course with the id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO: send a message if the course is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// get the id from the url
	params := mux.Vars(r)

	// loop through the courses to find the course with the id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted")
			return
		}
	}
}
