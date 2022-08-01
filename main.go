package main

import (
	"fmt"
	"sync"
	"time"
)

func generateNumbers(total int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("%d\n", total)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	wg.Add(1000000)
	for idx := 1; idx <= 1000000; idx++ {
		go generateNumbers(idx, &wg)
	}

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
	elapsed := time.Since(start)
	fmt.Printf("page took %d", elapsed.Milliseconds())
}
