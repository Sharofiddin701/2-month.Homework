package main

import (
	"fmt"
	"strings"
)

func findLongestWord(sentence string, resultChan chan int) {
	words := strings.Fields(sentence)
	longest := 0

	for _, word := range words {
		wordLength := len(word)
		if wordLength > longest {
			longest = wordLength
		}
	}

	resultChan <- longest
}

func main() {

	sentence := "Bir vaqtning o'zida berilgan satrdagi eng uzun so'zning uzunligini topish uchun goroutinlar va kanallardan foydalanadigan dastur yarating"

	resultChan := make(chan int)

	go findLongestWord(sentence, resultChan)

	longestWordLength := <-resultChan

	fmt.Printf("Eng uzun so'z uzunligi: %d\n", longestWordLength)
}
