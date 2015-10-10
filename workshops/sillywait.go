package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("hello from goroutine!")
	}()

	time.Sleep(time.Second * 1)
}
