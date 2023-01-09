package main

import (
	"fmt"
	"sync"
	"time"
)

func say(msg string) {
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup

	now := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			say("Hello")
		}()
	}
	wg.Wait()

	fmt.Println("Total time taken:", time.Since(now))
}
