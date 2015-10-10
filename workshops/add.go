package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Go hackers!")

	go func() {
		fmt.Println("hello from goroutine!")
	}()
}
