package main

import (
	"fmt"
)

func findLargest(numbers []int, result chan int) {
	if len(numbers) == 0 {
		result <- 0
		return
	}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	result <- max
}

func main() {

	result := make(chan int)

	numbers := []int{4, 23, 9, 12, 34, 8, 67, 55, 3, 1}

	go findLargest(numbers, result)
	largest := <-result
	fmt.Printf("the largest number %d\n", largest)
}
