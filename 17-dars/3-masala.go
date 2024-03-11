package main

import (
	"fmt"
)

func fakt(ch chan int) {
	S := 1
	var a int
	fmt.Print("Qiymat:")
	fmt.Scan(&a)

	go func() {
		for i := 1; i <= a; i++ {
			ch <- i
		}
		defer close(ch)
	}()

	for num := range ch {
		S *= num
	}
	fmt.Println("Fakto:", S)
}
func main() {

	ch := make(chan int)
	fakt(ch)
}
