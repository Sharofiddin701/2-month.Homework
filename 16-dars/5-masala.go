package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendRandomInt(ch chan<- int) {

	rand.Seed(time.Now().UnixNano())

	ch <- rand.Intn(100)
}

func main() {

	ch := make(chan int)

	go sendRandomInt(ch)

	num := <-ch
	fmt.Println("Received:", num)

	time.Sleep(time.Second)
}
