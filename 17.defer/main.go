package main

import "fmt"

/*
defer keyword takes the statement next to it just before the end of fn declaration, right before return
if there are more than one defer statements - last in first out. last defer is executed first and first, last
*/

/*
COMPLETE OUTPUT

Welcome to defer in Golang
World
World
in defer
4 3 2 1 0 Vani
Hello
Hello
*/
func main() {
	fmt.Println("Welcome to defer in Golang")
	defer fmt.Println("Hello")
	fmt.Println("World")
	// Output
	/*
		Welcome to defer in Golang
		World
		Hello
	*/

	defer fmt.Println("Hello")
	defer fmt.Println("Vani")
	fmt.Println("World")
	// Output
	/*
		World
		Vani
		Hello
	*/

	myDefer()
}

func myDefer() {
	fmt.Println("in defer")
	for i := range 5 {
		defer fmt.Print(i, " ")
	}
}
