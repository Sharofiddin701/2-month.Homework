package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(numbers chan<- int, count int) {
	defer close(numbers)
	for i := 0; i < count; i++ {
		numbers <- rand.Intn(100) 
	}
}

func consumer(numbers <-chan int) {
	for num := range numbers {
		fmt.Println("Chop etilgan raqam:", num)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())


	numbers := make(chan int)

	go producer(numbers, 10)
	go consumer(numbers)

	fmt.Println("Dastur tugaguncha kutamiz...")
	time.Sleep(3 * time.Second)
}
