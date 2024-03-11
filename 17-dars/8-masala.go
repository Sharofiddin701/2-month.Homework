package main

import (
	"fmt"
)

func element1(ch chan int) {
	a := []int{1, 24, 3, 4, 55, 6}
	var S int
	var l, k float64

	go func() {
		for _, num := range a {
			ch <- num
		}
		close(ch)
	}()

	for num := range ch {
		k++
		S += num
	}
	l = float64(S) / k

	fmt.Println("O'rtacha qiymati:", l)
}

func main() {
	ch := make(chan int)
	element1(ch)
}
