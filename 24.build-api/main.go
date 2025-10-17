package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for course - should be in a file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"` // use - to not show price in json
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB

var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to APIs in Golang")
	r := mux.NewRouter()

	// seeding data
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJs", CoursePrice: 299, Author: &Author{
		Fullname: "Vani Sharma", Website: "vani.go.dev",
	}})
	courses = append(courses, Course{CourseId: "2", CourseName: "NextJs", CoursePrice: 399, Author: &Author{
		Fullname: "Tripti", Website: "tripti.go.dev",
	}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")

	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	r.HandleFunc("/courses", deleteAllCourses).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))

}

// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API in Golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) // takes json and returns
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one courses")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses and find matching id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course with given id - " + params["id"])
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Empty JSON received")
		return
	}

	// checking if course already exists
	for _, crse := range courses {
		if crse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course already present in DB")
			return
		}
	}

	// generate a uniue id
	course.CourseId = strconv.Itoa(rand.IntN(100))

	// append course to courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	json.NewEncoder(w).Encode("Added a course successfully! ")
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop through the id -> remove and add with my id
	for key, course := range courses {
		if course.CourseId == params["id"] {
			// removing index key from the slice
			courses = append(courses[:key], courses[key+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course.CourseId + " updated")
			return
		}
	}
	json.NewEncoder(w).Encode(params["id"] + " not found!")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for ind, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:ind], courses[ind+1:]...)
			json.NewEncoder(w).Encode(params["id"] + " deleted")
			return
		}
	}

	json.NewEncoder(w).Encode(params["id"] + " not found for deletion")
}

func deleteAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete all courses")
	w.Header().Set("Content-Type", "application/json")

	courses = nil
	json.NewEncoder(w).Encode("All courses deleted")
}
