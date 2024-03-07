package main

import (
	"fmt"
	"time"
)

func worker(id int, done chan struct{}) {
	defer func() {
		done <- struct{}{}
	}()

	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	numWorkers := 10
	done := make(chan struct{})

	for i := 0; i < numWorkers; i++ {
		go worker(i, done)
	}

	timeout := time.After(100 * time.Second)
	for i := 0; i < numWorkers; i++ {
		select {
		case <-done:
			fmt.Println("Received signal from a worker.That worker is done")
		case <-timeout:
			fmt.Println("Timeout reached.Exiting....")
			return
		}
	}
	fmt.Println("All workers finished")
}
