package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go printMessage(ch)

	ch <- "Assalomu alaykum!"

	time.Sleep(time.Second)
}
func printMessage(ch <-chan string) {

	msg := <-ch
	fmt.Println("Goroutine: ", msg)

	fmt.Println("Dastur tugadi.")
}
