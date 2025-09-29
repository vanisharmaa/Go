package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to Math in Golang")

	var num1 int = 2
	var num2 float32 = 5.2

	// fmt.Println("Adding num1 and num2: ", num1+num2) // not allowed
	fmt.Println("Adding num1 and num2: ", num1+int(num2))

	// all three will return diff numbers
	fmt.Println("random number: ", rand.Intn(7)) // [0, 7) 0-6 numbers returned
	fmt.Println("random number: ", rand.Intn(7))
	fmt.Println("random number: ", rand.Intn(7))

	// NOTE:
	// A seed is the starting point for pseudo-random numbers.
	// Before Go 1.20, we had to call rand.Seed() to get different numbers each run.
	// Now the global generator is auto-seeded, so we don’t need it.
	// Use rand.New(rand.NewSource(seed)) for reproducible sequences.

	// rand.Seed(42) // deprecated now - time.Now().UnixNano() for random numbers
	// how to achieve seed now
	src := rand.NewSource(42)
	r := rand.New(src)
	fmt.Println("seed random number (always same): ", r.Intn(7))
}
