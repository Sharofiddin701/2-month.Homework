package main

import (
	"fmt"
	"strings"
)

func countWords(text string, resultChan chan<- int) {
	words := strings.Fields(text)
	wordCount := len(words)
	resultChan <- wordCount
}

func main() {
	text := "Bu birinchi misol matni. Ikkinchi misol matnida, so'zlarning sonini hisoblash uchun goroutinlar va kanallardan foydalaniladi."

	resultChan := make(chan int)

	go countWords(text, resultChan)

	wordCount := <-resultChan

	fmt.Println("Matndagi so'zlar soni:", wordCount)
}
