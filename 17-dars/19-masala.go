package main

import (
	"bufio"
	"fmt"
	"os"
)

func processFile(filename string, results chan int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Xatolik: Faylni ochishdagi xatolik:", err)
		results <- 0
		return
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Xatolik: Fayldagi satrlarni o'qishdagi xatolik:", err)
		results <- 0
		return
	}

	results <- lineCount
}

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}

	results := make(chan int, len(files))

	for _, filename := range files {
		go processFile(filename, results)
	}

	totalLines := 0
	for i := 0; i < len(files); i++ {
		totalLines += <-results
	}

	fmt.Println("Jami satrlar soni:", totalLines)
}
