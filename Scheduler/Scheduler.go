package scheduler

import (
	"fmt"
	"time"
)

func task(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Task %d - Iteration %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go task(1)
	go task(2)
	time.Sleep(1 * time.Second)
}
