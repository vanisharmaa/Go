package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"` // renames Name to courseName in jsons
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // removes the field from response
	Tags     []string `json:"tags,omitempty"` // removes empty (nil or null) fields from response. do not add space before omitempty
}

func main() {
	fmt.Println("Welcome to JSON in Golang")
	// Encoding the JSON
	EncodeJson()
}

func EncodeJson() {
	courses := []course{
		{"reactjs bootcamp", 299, "udemy.com", "abc123", []string{"web-dev", "js", "frontend"}},
		{"nextjs bootcamp", 499, "udemy.com", "def123", []string{"web-dev", "js", "fullstack"}},
		{"angularjs bootcamp", 299, "udemy.com", "ghi123", nil},
		{"go bootcamp", 499, "udemy.com", "jkl123", []string{"backend", "frontend", "server"}},
		{"cpp bootcamp", 399, "udemy.com", "mno123", []string{"coding", "c", "language"}},
	}

	// package this data as JSON data

	// finalJSON, err := json.Marshal(courses)
	finalJSON, err := json.MarshalIndent(courses, "*", "\t") // result is indented based on tab and every line begins with a *

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON) // nil is read as null in json
}
