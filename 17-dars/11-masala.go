package main

import (
	"fmt"
)

func kv() {
	ch := make(chan int)
	s := 0

	go func() {
		for i := 1; i <= 1000; i++ {
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		s += num * num
	}

	fmt.Println("Kvadratlari yig'indisi:", s)
}

func main() {
	kv()
}
