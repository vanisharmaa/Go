package main

import "fmt"

type User struct { // begin struct name with capital coz public thingy - similarly with keys
	Id     int
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in Golang!")

	// - no inheritance
	// - no super or parent class

	vani := User{
		1,
		"Vani",
		"vani@go.dev",
		true,
		24,
	}

	fmt.Println("vani: ", vani)                          // {1 vani vani@go.dev true 24}
	fmt.Printf("User Vani's details are: \n%+v\n", vani) // {Id:1 Name:vani Email:vani@go.dev Status:true Age:24}
	fmt.Printf("Name is: %s\n", vani.Name)
}
