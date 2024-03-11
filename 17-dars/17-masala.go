package main

import (
	"fmt"
)

func kv() {
	ch := make(chan int)
	s := 0

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		if num%2 == 0 {
			s += num
		}
	}

	fmt.Println("Juft sonlar yigindisi:", s)
}

func main() {
	kv()
}
