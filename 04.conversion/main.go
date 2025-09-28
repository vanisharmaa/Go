package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1-5")

	reader := bufio.NewReader(os.Stdin)
	rating, _ := reader.ReadString('\n')
	// The input currently is "4\n", using below function - it'll only be "4"
	rating = strings.TrimSpace(rating)

	fmt.Println("Thanks for your rating - ", rating)
	fmt.Printf("Type of rating %T\n", rating)

	// This function converts string to float - the second arguments is bits (32 or 64 for float)
	// The bitSize argument (e.g., 32 or 64) only affects input validation (the allowed range).
	// The return type is always int64 for integers, or float64 for floats, regardless of bitSize.

	numRating, err := strconv.ParseFloat(rating, 32)
	numRatingInt, _ := strconv.ParseInt(rating, 10, 32) // (string, base, bits) - base can be 10 for decimal. 8 for octal, 16 for hexadecimal and so on

	if err != nil {
		fmt.Println("ERROR! - ", err)
	} else {
		fmt.Println("Added 1 to User Rating - ", numRating+1)
		fmt.Printf("Type of rating - %T\n", numRating)

		fmt.Println("Added 1 to User Rating in INT - ", numRatingInt+1)
		fmt.Printf("Type of rating - %T\n", numRatingInt)

	}
}
