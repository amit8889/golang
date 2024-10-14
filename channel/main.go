package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a new instance of the struct
	//wait group
	var wg sync.WaitGroup
	start := time.Now()
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go printVar(i, &wg)
	}
	end := time.Since(start)
	wg.Wait()
	fmt.Println()
	fmt.Println("===", end)

}

func printVar(i int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Print(i)
}
