package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello, Go hackers!")

	// Create a wait group.
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("hello from goroutine!")
		wg.Done()
	}()

	// Wait until everyone finishes.
	wg.Wait()
}
