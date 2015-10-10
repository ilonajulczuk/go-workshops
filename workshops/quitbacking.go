package main

import (
	"fmt"
	"time"
)

func bakeCakes(quit <-chan struct{}, output chan<- string) {
	for {
		select {
		case output <- "Cake":
		case <-quit:
			close(output)
			return
		}
	}
}

func main() {
	quit := make(chan struct{}, 1)
	cakes := make(chan string, 1)

	go bakeCakes(quit, cakes)

	time.AfterFunc(5*time.Second, func() {
		// START OMIT
		fmt.Println("It's time to close!")
		fmt.Println("Closing bakery from another goroutine.")
		select {
		case quit <- struct{}{}:
		default:
		}
		// END OMIT
	})

	var i int
	for i < 10 {
		time.Sleep(time.Second)
		_, ok := <-cakes
		if !ok {
			fmt.Println("Bakery was closed!")
			return
		}
		i++
		fmt.Printf("Consumed %d cakes so far.\n", i)
	}

	select {
	case quit <- struct{}{}:
	default:
	}
	fmt.Printf("Consumed %d cakes.\n", i)
}
