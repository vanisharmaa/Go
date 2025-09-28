package main

import "fmt"

// variable := "can i do this?" // no -> error
var variable = "can i do this?" // yes -> global variable
// fmt.Println("global var: ", variable) // not allowed

const Pi = 3.1412 // Public variable - capital first letter marks public thingies

func main() {
	var username string = "Vani"
	fmt.Println("User Name - ", username)
	fmt.Printf("Variable is of type - %T\n", username)

	var isLoggedIn bool = true
	fmt.Println("Is User Logged In - ", isLoggedIn)
	fmt.Printf("Variable is of type - %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println("a small value - ", smallVal)
	fmt.Printf("Variable is of type - %T\n", smallVal)

	var smallFloat float64 = 255.873857298284983 // float64 is more precise - more digits after decimal, float32 has lesser digits after decimal
	fmt.Println("a small float value - ", smallFloat)
	fmt.Printf("Variable is of type - %T\n", smallFloat)

	// default values and some aliases
	var anotherInt int
	var anotherString string
	fmt.Printf("another int: %d\nanother string: %s\n", anotherInt, anotherString) // 0, ""

	// implicit type
	var google = "https://www.google.com/"
	fmt.Printf("variable website: %s of type %T\n", google, google)

	// no var variable
	yahoo := "https://www.yahoo.com/"
	fmt.Printf("variable yahoo: %s of type %T\n", yahoo, yahoo)

	numUsers := 30000.00
	fmt.Printf("variable numUsers: %f of type %T\n", numUsers, numUsers)

	// global variable
	fmt.Println("global var: ", variable)

	// constants
	const pi float32 = 3.14
	fmt.Println("pi = ", pi)
	fmt.Println("Pi = ", Pi)
	fmt.Printf("type of pi is %T and Pi is %T\n", pi, Pi) // float32 (defined), float64 (implicit type)

}
