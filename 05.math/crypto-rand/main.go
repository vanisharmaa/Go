package main

import (
	"crypto/rand" // Secure random number generator
	"fmt"         // For printing output
	"math/big"    // Arbitrary-precision arithmetic (big integers)
)

func main() {
	// Generate a cryptographically secure random integer in [0, 10)
	// rand.Int uses crypto/rand.Reader as the source of randomness
	// big.NewInt(10) sets the upper bound (exclusive)
	r, _ := rand.Int(rand.Reader, big.NewInt(10))

	// Print the generated random number to the output pane
	fmt.Println("crypto r: ", r)
}
