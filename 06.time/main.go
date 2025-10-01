package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println("presentTime: ", presentTime)
	fmt.Println(presentTime.Format("02 Jan, 2006 Monday 03:04:05 PM IST"))
	currentDate := presentTime.Format("02 Jan, 2006")
	fmt.Printf("current date type is %T\n", currentDate)

	createdDate := time.Date(2001, time.September, 10, 18, 36, 0, 0, time.Local)
	prettyCreatedDate := createdDate.Format("02 Jan, 2006 - Monday, 03:04 PM IST")
	fmt.Println("created date: ", createdDate)
	fmt.Println("pretty created date: ", prettyCreatedDate)
}
