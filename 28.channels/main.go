package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")
	wg := &sync.WaitGroup{}
	myChannel := make(chan int, 2) // type of value in channel and number of values in channel

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) { // <- on the left means read only channel
		fmt.Println(<-ch)
		fmt.Println(<-ch) // read as many times as values are set
		// close(ch) !error
		wg.Done()
	}(myChannel, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) { // <- on the right means write only channel
		ch <- 5
		ch <- 6
		close(ch)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
