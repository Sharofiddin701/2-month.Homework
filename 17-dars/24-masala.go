package main

import (
	"fmt"
	"sync"
)

func main() {
	const numWorkers = 5
	const numTasks = 10

	tasks := make(chan int, numTasks)

	results := make(chan int, numTasks)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(tasks, results, &wg)
	}

	for i := 0; i < numTasks; i++ {
		tasks <- i
	}

	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("Natija:", res)
	}
}
func worker(tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {

		result := task * 2

		results <- result
	}
}
