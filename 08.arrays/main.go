package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Banana"

	fmt.Println(len(fruitList), " - fruitList: ", fruitList)

	var vegList = [3]string{
		"Mushrooms",
		"Bell-Peppers",
		"Potatoes",
	}
	fmt.Println(len(vegList), " - vegList: ", vegList)
}
