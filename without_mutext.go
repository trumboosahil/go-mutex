package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

func main() {
	counter := &Counter{}

	// Number of goroutines
	const numGoroutines = 100

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter.Value)
}
