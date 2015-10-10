package main

import (
	"fmt"
	"sync"
)

type Knight struct {
	mu     sync.Mutex
	health int
}

func main() {
	knight := &Knight{health: 100}
	var wg sync.WaitGroup

	attack := func(k *Knight) {
		k.mu.Lock()
		defer k.mu.Unlock()
		k.health = k.health - 10
		wg.Done()
	}

	heal := func(k *Knight) {
		k.mu.Lock()
		defer k.mu.Unlock()
		k.health = k.health + 10
		wg.Done()
	}

	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go attack(knight)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go heal(knight)
		}
		wg.Done()
	}()
	wg.Wait()
	knight.mu.Lock()
	fmt.Printf("%#v", knight)
	knight.mu.Unlock()
}
