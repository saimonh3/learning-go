package main

import (
	"fmt"
	"sync"
)

func main() {
	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex

	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		mu.Unlock()
	}

	withdraw := func (amount int) {
		mu.Lock()
		balance -= amount
		mu.Unlock()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdraw(1)
		}()
	}

	wg.Wait()
	fmt.Println(balance)
}