package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in Golang!")

	days := make([]string, 7)
	days[0] = "mon"
	days[1] = "tues"
	days[2] = "wed"
	days[3] = "thurs"
	days[4] = "fri"
	days[5] = "sat"
	days[6] = "sun"

	// for loop
	for d := 0; d < len(days); d++ {
		fmt.Println("day ", d+1, " = ", days[d])
	}

	// for - range
	for i := range days { // i returns index only - 0, 1, 2...
		fmt.Println(i+1, " -> ", days[i])
	}

	// for - range (i is index, day is the actual ith index of slice)
	for i, day := range days {
		fmt.Println(i+1, " --- ", day)
	}

	/*******************************************************************/

	// while-like for loop
	val := 1
	for val <= 10 {
		if val == 5 {
			val++
			continue
		} else if val == 8 {
			break
		} else if val == 7 {
			goto label
		}

		fmt.Println(val)
		val++
	}

label:
	fmt.Println("Jumped to label")
}
