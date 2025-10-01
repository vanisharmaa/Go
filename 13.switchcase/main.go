package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to Switch Case in Golang!")

	dieNum := rand.Intn(6)
	dieNum++
	fmt.Println("die number: ", dieNum)
	switch dieNum {
	case 1:
		fmt.Println("You may now begin your game!")
	case 2, 3, 4, 5:
		fmt.Printf("Move %v places\n", dieNum)
	case 6:
		fmt.Println("Move 6 places & pls roll the die again!")
	default:
		fmt.Println("Wtf was that bruh?!")
	}

	// fallthrough

	num := rand.Intn(10)
	fmt.Println("num = ", num)
	switch num {
	case 0, 1, 2, 3, 4, 5, 6, 7:
		fallthrough
	case 8:
		fmt.Println("you got a number <= 8")
	case 9:
		fmt.Println("WooHooooo 9!")
	default:
		fmt.Println("Umm???")
	}

}
