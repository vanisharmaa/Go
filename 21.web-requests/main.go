// Staer the webserver before running this file
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url string = "http://localhost:8000/"

func main() {
	fmt.Println("Welcome to web requests in Golang")

	// const getUrl string = url + "get/"

	// GET/ requests
	/* PerformGetRequests(getUrl)
	PerformGetRequestsMethod2(getUrl) */

	const postUrl string = url + "post/"

	PerformPostJsonRequest(postUrl)

}

func PerformGetRequests(myUrl string) {
	// "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	fmt.Printf("type of response %T\n", content)
	fmt.Println("response: ", string(content))
}

func PerformGetRequestsMethod2(myUrl string) {
	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content) // same as content length

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println("response: ", responseString.String()) // response
}

func PerformPostJsonRequest(myUrl string) {
	// fake json payload
	// create any type of data you want using strings.NewReader
	// this creates json
	requestBody := strings.NewReader(`
		{
			"name": "vani",
			"age": "24",
			"country": "India",
			"profession": "Software Engineer"
		}
	`)

	fmt.Printf("Type of request body = %T\n", requestBody)
	response, err := http.Post(myUrl, "Application/json", requestBody)

	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("content: ", string(content))
}
