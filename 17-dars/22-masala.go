package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0
	numGoroutines := 5

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Final counter value:", counter)
}
