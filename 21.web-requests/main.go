// Staer the webserver before running this file
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const baseUrl string = "http://localhost:8000/"

func main() {
	fmt.Println("Welcome to web requests in Golang")

	const getUrl string = baseUrl + "get/"

	// GET/ requests
	PerformGetRequests(getUrl)
	PerformGetRequestsMethod2(getUrl)

	const postJsonUrl string = baseUrl + "post/"

	// POST/ json request
	PerformPostJsonRequest(postJsonUrl)

	const postFormUrl string = baseUrl + "postform/"

	// POST/ form request
	PerformPostFormRequest(postFormUrl)

}

func PerformGetRequests(myUrl string) {
	// "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("GET/ Status code: ", response.StatusCode)
	fmt.Println("GET/ Content Length: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	fmt.Printf("GET/ type of response %T\n", content)
	fmt.Println("GET/ response: ", string(content))
}

func PerformGetRequestsMethod2(myUrl string) {
	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("2. GET/ Status code: ", response.StatusCode)
	fmt.Println("2. GET/ Content Length: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content) // same as content length

	fmt.Println("2. GET/ ByteCount is: ", byteCount)
	fmt.Println("2. GET/ response: ", responseString.String()) // response
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

	fmt.Printf("Json POST/ Type of request body = %T\n", requestBody)
	response, err := http.Post(myUrl, "Application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Json POST/ content: ", string(content))
}

func PerformPostFormRequest(myUrl string) {
	// form data
	data := url.Values{}
	data.Add("fname", "vani")
	data.Add("lname", "sharma")
	data.Add("age", "24")
	data.Add("country", "india")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	byteRes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Form POST/ post form response: ", string(byteRes))
}
