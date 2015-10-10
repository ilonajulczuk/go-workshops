package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Food struct {
	Name string
}

type Person struct {
	Name string
}

func (p Person) Eat(food Food) {
	fmt.Printf("%s is enjoying some %s.\n", p.Name, food.Name)
}

func main() {
	amount := 5
	foods := []Food{
		{"icecream"},
		{"salad"},
		{"cake"},
		{"fish"},
		{"potatoes"},
	}
	var setsOfFood []chan Food
	for _, food := range foods {
		newSet := make(chan Food, amount)
		for i := 0; i < amount; i++ {
			newSet <- food
		}
		setsOfFood = append(setsOfFood, newSet)
	}

	fmt.Println("Bon appétit!")

	var wg sync.WaitGroup
	for _, person := range []Person{{"Alice"}, {"Bob"}, {"Eve"}, {"John"}, {"Philip"}} {
		wg.Add(1)
		go func(person Person) {
			perm := rand.Perm(amount)
			for _, i := range perm {
				time.Sleep(time.Millisecond * time.Duration(rand.Int()%100))
				dish := <-setsOfFood[i]
				person.Eat(dish)
			}
			wg.Done()
		}(person)
	}
	wg.Wait()

	fmt.Println("That was delicious!")
}
