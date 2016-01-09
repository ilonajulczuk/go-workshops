package main

import (
	"fmt"
)

func main() {
	itemNo := make(chan int)
	coins := make(chan int)
	itemCh := make(chan int)

	go func() {
		fmt.Println("Vending machine is starting")
		for {
			item := <-itemNo
			fmt.Println("Item: ", item, "choosen")
			itemValue := 9
			var coinsValue int
			for coinsValue < itemValue {
				coinsValue += <-coins
			}
			itemCh <- item
		}
	}()
	itemNo <- 12
	coins <- 5
	coins <- 2
	coins <- 2
	rec := <-itemCh
	fmt.Println("Item: ", rec, "recived")
}
