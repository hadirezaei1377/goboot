package sharedvariables

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	data := []int{1, 2, 3, 4, 5}

	results := make(chan int)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, v := range data {
			results <- v * 2 //send results to channel
		}
	}()

	go func() {
		defer wg.Done()
		for _, v := range data {
			results <- v * 3 //send results to channel
		}
	}()

	go func() {
		wg.Wait()
		close(results) // goroutines is done
	}()

	for result := range results {
		fmt.Println(result)
	}
}
