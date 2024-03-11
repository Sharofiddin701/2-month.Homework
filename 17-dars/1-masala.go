package main

import (
	"fmt"
)

func exer(ch chan int) {
	defer close(ch)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	for num := range ch {
		fmt.Println("Qiymat:", num)
	}
}

func main() {
	ch := make(chan int)
	exer(ch)
}
