package main

import (
	"fmt"
)

func summ(ch chan int) {
	S := 0

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		defer close(ch)
	}()

	for num := range ch {
		S += num
	}
	fmt.Println("Summasi:", S)
}
func main() {
	ch := make(chan int)
	summ(ch)
}
