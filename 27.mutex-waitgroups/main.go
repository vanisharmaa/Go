package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition via goroutines in golang")

	var score = []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	wg.Add(4) // or we can do wg.Add(1) before every goroutine
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.RLock()
		fmt.Println("score in goroutine - ", score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println("score - ", score)
}
