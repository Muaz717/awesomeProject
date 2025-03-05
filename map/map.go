package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := make(map[int]int)
	mu := &sync.Mutex{}

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("Goroutine", i)
			mu.Lock()
			m[i] = i + 1
			mu.Unlock()
		}()
	}

	time.Sleep(1 * time.Second)

	fmt.Println(m)

}
