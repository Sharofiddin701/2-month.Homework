package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetchURL(url string, a, b <-chan string, statusChan chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		statusChan <- fmt.Sprintf("Error fetching %s: %s", url, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		statusChan <- fmt.Sprintf("Error reading response from %s: %s", url, err)
		return
	}

	statusChan <- fmt.Sprintf("Fetched %s: %d bytes", url, len(body))

	// Read requests from channels a and b
	select {
	case reqA := <-a:
		fmt.Printf("Received request from channel A: %s\n", reqA)
	case reqB := <-b:
		fmt.Printf("Received request from channel B: %s\n", reqB)
	default:
		// No requests received
	}
}

func main() {
	urls := []string{
		"http://example.com",
		"https://www.google.com",
		"https://www.github.com",
		"https://www.wikipedia.org",
	}
	statusChan := make(chan string)
	channelA := make(chan string)
	channelB := make(chan string)

	timeout := time.After(3 * time.Second) // Set a timeout of 3 seconds

	for _, url := range urls {
		go fetchURL(url, channelA, channelB, statusChan)
	}

	for i := 0; i < len(urls); i++ {
		select {
		case status := <-statusChan:
			fmt.Println(status)
		case <-timeout:
			fmt.Println("Timeout reached. Exiting...")
			return
		}
	}

	// Closing channels to signal no more requests
	close(channelA)
	close(channelB)
}
