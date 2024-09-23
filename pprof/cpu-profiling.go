package pprof

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func heavyOperation(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			_ = fmt.Sprintf("Number: %d", i*j)
		}
	}
}

func main() {
	// start pprof on 6060 port
	go func() {
		fmt.Println("Starting pprof on :6060")
		http.ListenAndServe("localhost:6060", nil)
	}()

	fmt.Println("Running heavy operation...")
	heavyOperation(5000)

	// stop application for profiling
	time.Sleep(10 * time.Minute)
}
