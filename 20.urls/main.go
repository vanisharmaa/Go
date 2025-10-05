package main

import (
	"fmt"
	"net/url"
)

var myUrl string = "https://vani.dev:3000/learn?coursename=reactjs&paymentid=b73gfhyj832rfiqu38"

// "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

func main() {
	fmt.Println("Welcome to handling URLs in Golang")
	fmt.Println(myUrl)

	// parsing the url
	res, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println("result scheme: ", res.Scheme)      // https
	fmt.Println("result host: ", res.Host)          // vani.dev:3000
	fmt.Println("result path: ", res.Path)          // /learn
	fmt.Println("result raw query: ", res.RawQuery) // coursename=reactjs&paymentid=b73gfhyj832rfiqu38
	fmt.Println("result port: ", res.Port())        // 3000

	qParams := res.Query()
	fmt.Printf("The type of qParams is %T\n", qParams) // url.Values (key value pair)
	fmt.Println("query params: ", qParams)
	fmt.Println(qParams["coursename"], qParams["paymentid"])

	for key, val := range qParams {
		fmt.Println(key, " : ", val)
	}

	partsOfUrl := &url.URL{ // ! Pass reference only
		Scheme:   "https",
		Host:     "vani.dev",
		Path:     "/learn",
		RawQuery: "user=vani",
	}

	url2 := partsOfUrl.String()

	fmt.Println("url 2 : ", url2) // https://vani.dev/learn?user=vani

}
