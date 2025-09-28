package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) // creates a buffer object
	fmt.Println("Please give us a rating")

	// reader.ReadString(('\n')) actually takes the input till \n is found

	// comma ok || comma error syntax
	// rating, err := reader.ReadString('\n') // go provides output/error as two variables - \n means take input till the user presses a newline character or enter
	// if you don't wanna use one of the variables, you can just use an underscore (_) in its place like
	rating, _ := reader.ReadString(('\n'))

	fmt.Println("Thanks for your rating, ", rating)
	fmt.Printf("Type of rating variable - %T\n", rating) // string
}
