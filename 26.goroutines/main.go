package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// 	fmt.Println("Welcome to goroutines in Golang")
// 	go greeter("hello")
// 	greeter("world")
// }

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(i+1, s)
// 	}
// 	fmt.Println()
// }

var wg sync.WaitGroup

var mut sync.Mutex

var db = []string{"test"}

func main() {
	websites := []string{
		"https://www.google.com",
		"https://go.dev",
		"https://www.youtube.com",
		"https://fb.com",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println("db - ", db)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d from %s", res.StatusCode, endpoint)

	mut.Lock()

	db = append(db, endpoint)

	mut.Unlock()

	fmt.Println("")
}
