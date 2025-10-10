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
	// EncodeJson()

	DecodeJson()
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

func DecodeJson() {
	jsonData := []byte(`
		{
			"courseName": "go bootcamp",
			"Price": 499,
			"website": "udemy.com",
			"tags": [
					"backend",
					"frontend",
					"server"
			]
		}
	`)

	var course1 course

	isValid := json.Valid(jsonData)

	if isValid {
		fmt.Println("json is valid")
		json.Unmarshal(jsonData, &course1)
		fmt.Printf("%#v\n ", course1)
		fmt.Printf("type of course1 is %T and json data is of type %T\n", course1, jsonData)
	} else {
		fmt.Println("JSON NOT VALID!")
	}

	// key-value pair

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Println(myOnlineData)

	for k, v := range myOnlineData {
		fmt.Println(k, " -> ", v)
		fmt.Printf("Type of value is %T\n", v)

	}
}
