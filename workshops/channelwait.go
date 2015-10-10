package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("hello from goroutine!")
		ch <- "done"
	}()

	<-ch
}
