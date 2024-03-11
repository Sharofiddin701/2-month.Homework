package main

import (
	"fmt"
)

func satr(ch chan int) {
	var a string
	var s string
	fmt.Print("So'zni kiriting:")
	fmt.Scan(&a)

	go func() {
		defer close(ch)
		for i := len(a) - 1; i >= 0; i-- {
			ch <- i
		}
	}()

	for num := range ch {
		s += string(a[num])
	}
	fmt.Println("Teskari so'z:", s)
}

func main() {
	ch := make(chan int)
	satr(ch)
}
