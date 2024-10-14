package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("=======main function=====")

	var num int32 = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	//var rmux sync.RWMutex
	//rmux.RLock()
	//rmux.Lock()
	fmt.Println("=====num====>", num)
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		// go inc(&num, &wg, &mu)
		go inc(&num, &wg, &mu)
	}
	wg.Wait()
	fmt.Println("====final=num====>", num)
}
func inc(n *int32, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer func() {
		mu.Unlock()
		wg.Done()
	}()
	mu.Lock()
	*n++

}

// func inc(n *int32, wg *sync.WaitGroup, mu *sync.RWMutex) {
// 	defer wg.Done()
// 	mu.RLock()
// 	*n++
// 	mu.RUnlock()
// }
