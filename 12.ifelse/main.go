package main

import "fmt"

func main() {
	fmt.Println("If Else in Golang!")

	if 2 > 10 {
		fmt.Println("If")
	} else if 2 == 10 {
		fmt.Println("Else if")
	} else {
		fmt.Println("Else")
	}

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// Declare a variable and use in if

	if num := 10; num < 10 {
		fmt.Println(num, " < 10")
	} else if num == 10 {
		fmt.Println(num, " = 10")
	} else {
		fmt.Println(num, " > 10")
	}
}
