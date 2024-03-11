package main

import (
	"fmt"
)

func isPrime(num int, primeChan chan<- bool) {
	if num <= 1 {
		primeChan <- false
		return
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			primeChan <- false
			return
		}
	}
	primeChan <- true
}

func findAllPrimesInRange(start, end int, primeChan chan<- int) {
	for i := start; i <= end; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime && i > 1 {
			primeChan <- i
		}
	}
	close(primeChan)
}

func main() {
	var start, end int
	fmt.Print("Diapazonni kiriting: ")
	fmt.Scan(&start, &end)

	primeChan := make(chan int)

	go findAllPrimesInRange(start, end, primeChan)

	fmt.Println("Diapazondagi tub sonlar:")
	for prime := range primeChan {
		fmt.Println(prime)
	}
}
