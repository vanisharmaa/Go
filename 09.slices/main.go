package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices!")
	var fruitList = []string{
		"apples",
		"mangoes",
		"chikoos",
	}
	fmt.Printf("type of fruitList is %T\n", fruitList) // []string
	fmt.Println(fruitList, fruitList[2])               // fruitList[4] will give an error

	fruitList = append(fruitList, "litchis")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:]) // slices fruitList - removes 0th element
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // slices fruitList - 1st and 2nd element only
	fmt.Println(fruitList)

	fruitList = append(fruitList[:1]) // slices fruitList - only keeps the first element
	fmt.Println(fruitList)

	/************************************************************************************/

	highScores := make([]int, 4)
	fmt.Println(highScores) // [0 0 0 0]

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 // error

	highScores = append(highScores, 555, 666, 777) // ok
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println("sorted slice: ", highScores)
	fmt.Println("is slice sorted: ", sort.IntsAreSorted(highScores))

	/************************************************************************************/
	// remove values from a slice

	var courses = []string{
		"reactjs",
		"javascript",
		"swift",
		"python",
		"ruby",
	}

	fmt.Println("courses: ", courses)

	var indexToRemove = 2

	courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...) // spread operator
	fmt.Println("after remocing first index: ", courses)

}
