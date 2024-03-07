package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendRandomNumber() (chan int, <-chan string) {
	randomCh := make(chan int)
	readableCh := make(chan string)
	go func() {
		defer close(randomCh)
		defer close(readableCh)
		for {
			select {
			case randomCh <- rand.Intn(100):
				readableCh <- "Number sent"
			case <-time.After(time.Second):
				readableCh <- "Timeout occurred"
				return
			}
		}
	}()
	return randomCh, readableCh
}
func main() {
	randomCh, readableCh := sendRandomNumber()
	for {
		select {
		case num := <-randomCh:
			fmt.Println("Recieivid random number:", num)
		case status := <-readableCh:
			fmt.Println("Status", status)
		}
	}
}
