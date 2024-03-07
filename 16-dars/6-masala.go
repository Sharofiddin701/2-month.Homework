package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumbers(channel chan int, count int) {
	defer close(channel)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		channel <- rand.Intn(100)
	}
}

func main() {
	channel := make(chan int)
	go randomNumbers(channel, 3)
	for num := range channel {
		fmt.Println("Random son:", num)
	}
}
