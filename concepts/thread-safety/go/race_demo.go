package main

import (
	"fmt"
	"sync"
	"time"
)

// This program demonstrates a Race Condition crashing the app.
// Run with: go run race_demo.go
func main() {
	fmt.Println("Starting unsafe concurrent map access...")

	// Shared Map
	m := make(map[int]int)

	// WaitGroup just waits for threads to finish
	var wg sync.WaitGroup

	// Spawning 5 "Writers"
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				// WRITE
				m[j] = j + id
				time.Sleep(time.Millisecond) // Simulate work
			}
		}(i)
	}

	// Spawning 5 "Readers"
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				// READ
				_ = m[j]
				time.Sleep(time.Millisecond)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Done! (If you see this, you got lucky. Usually it crashes)")
}
