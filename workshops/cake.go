package main

import (
	"fmt"
	"time"
)

func main() {
	chocolateCakeCh := make(chan string)
	tiramisuCh := make(chan string)
	go func() {
		for {
			time.Sleep(time.Second)
			chocolateCakeCh <- "Chocolate cake"
		}
	}()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			tiramisuCh <- "Tiramisu cake"
		}
	}()

	// START OMIT
	timeout := time.After(5 * time.Second)
	var cake string
	var done bool
	for !done {
		select {
		case cake = <-chocolateCakeCh:
			fmt.Printf("Omnomnom, delicious %s!\n", cake)
		case cake = <-tiramisuCh:
			fmt.Printf("Fantastic! Lift me up! %s!\n", cake)
		case <-timeout:
			fmt.Println("That was great, but I have to go home!")
			done = true
		}
	}
	// END OMIT
}
