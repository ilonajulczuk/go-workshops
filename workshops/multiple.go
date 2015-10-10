package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(42)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(number int) {
			time.Sleep(time.Millisecond * 10 * time.Duration(rand.Int()%50))
			fmt.Println("hello from goroutine!", number)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
