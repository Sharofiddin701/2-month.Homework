package main

import (
	"fmt"
	"strings"
)

func countVowelsAndConsonants(text string, resultChan chan<- map[string]int) {
	vowels := "aeiouAEIOU"
	vowelCount := 0
	consonantCount := 0

	for _, char := range text {
		if strings.ContainsRune(vowels, char) {
			vowelCount++
		} else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			consonantCount++
		}
	}
	result := map[string]int{
		"vowels":     vowelCount,
		"consonants": consonantCount,
	}

	resultChan <- result
}

func main() {
	text := "Sharofiddin"

	resultChan := make(chan map[string]int)

	go countVowelsAndConsonants(text, resultChan)

	result := <-resultChan

	fmt.Println("Unli harflar soni:", result["vowels"])
	fmt.Println("Undoshlar soni:", result["consonants"])
}
