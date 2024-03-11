package main

import (
	"fmt"
	"math"
)

func calculateSquareRoot(num int, resultChan chan<- float64) {
	squareRoot := math.Sqrt(float64(num))
	resultChan <- squareRoot
}
func calculateAndPrintSquareRoots(numbers []int, resultChan chan<- float64) {

	for _, num := range numbers {
		go calculateSquareRoot(num, resultChan)
	}

	for i := 0; i < len(numbers); i++ {
		squareRoot := <-resultChan
		fmt.Printf("Square root of %d: %.2f\n", numbers[i], squareRoot)
	}
}
func main() {
	numbers := []int{4, 9, 16, 25, 36}

	resultChan := make(chan float64, len(numbers))

	calculateAndPrintSquareRoots(numbers, resultChan)
}
