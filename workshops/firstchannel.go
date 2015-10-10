package main

import (
	"fmt"
)

func main() {
	// Create channel of strings.
	ch := make(chan string)

	// Send to channel.
	go func() { ch <- "ping" }()

	// Receive entry from channel and assign it to variable.
	message := <-ch
	fmt.Println("Message from channel:", message)
}
