package main

import (
	"fmt"
)

func fibonacci(n int, resultChan chan<- int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		resultChan <- a
		a, b = b, a+b
	}
	close(resultChan)
}

func main() {
	n := 10

	resultChan := make(chan int)

	go fibonacci(n, resultChan)

	fmt.Printf("Fibonacci sequence up to the %dth element:\n", n)
	for fib := range resultChan {
		fmt.Println(fib)
	}
}
