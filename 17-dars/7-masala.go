package main

import (
	"fmt"
)

func element(ch chan int) {
	a := []int{1, 24, 3, 4, 55, 6}

	go func() {
		for _, num := range a {
			ch <- num
		}
		close(ch)
	}()

	for num := range ch {
		S := 2 * num
		fmt.Println("yangilangan holati:",num,":", S)
	}
}
func main() {
	ch := make(chan int)
	element(ch)
}
