package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	errCh := make(chan error)

	numGoroutines := 5

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {

			err := doTask()

			if err != nil {
				errCh <- err
			}

			wg.Done()
		}()
	}

	go func() {

		for err := range errCh {
			fmt.Println("Error:", err)
		}
	}()
	wg.Wait()

	close(errCh)
}
func doTask() error {

	if randomError() {
		return errors.New("Task encountered an error")
	}
	return nil
}

func randomError() bool {

	return true
}
