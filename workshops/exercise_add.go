package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(42)
	for i := 0; i < 10; i++ {
		func(number int) {
			time.Sleep(time.Millisecond * 10 * time.Duration(rand.Int()%50))
			fmt.Println("hello from goroutine!", number)
		}(i)
	}
}
