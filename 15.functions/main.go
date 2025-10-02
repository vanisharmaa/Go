package main

import "fmt"

func main() { // func keyword for function
	fmt.Println("Welcome to functions in Golang!")
	greeter2()
	greeter()

	a, b := 5, 71
	fmt.Println(add(a, b))

	fmt.Println(proAdd(5, 6, 7, 9, 10))

	fmt.Println(addSub(10, 5))

}

// return multiple values
func addSub(a int, b int) (int, int) {
	return a + b, a - b
}

func proAdd(values ...int) int { // expecting multiple integer values
	fmt.Printf("type of values is %T\n", values) // []int = integer slices
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func add(a int, b int) int {
	return a + b
}

func greeter() {
	fmt.Println("Namaste from Golang!")
}

func greeter2() {
	fmt.Println("Hello from Golang!")
}
