package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers!") // Print welcome message

	var one int = 1 // Declare an integer variable 'one' and assign value 1

	var two *int // Declare a pointer to int 'two', default value is nil

	var three *int = &one // Declare a pointer 'three' and assign the address of 'one'

	one++ // Increment the value of 'one' by 1

	// Print values:
	// 'one' (current value),
	// 'two' (nil pointer),
	// '&one' (address of one = three)
	// 'three' (address of 'one'),
	// '*three' (value at the address pointed by 'three')
	fmt.Println(one, two, &one, three, *three)
}
