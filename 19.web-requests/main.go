package main

import (
	"fmt"
	"io"
	"net/http"
)

var url string = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

func main() {
	fmt.Println("Welcome to Web Requests in Golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)
	fmt.Println("response:= ", response)

	// always close an http connection yourself
	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("databytes in string: ", string(dataBytes))
}
