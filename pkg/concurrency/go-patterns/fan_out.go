package go_patterns

import (
	"fmt"
	"time"
)

func worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- process(j)
	}
}

func process(j int) int {
	time.Sleep(1 * time.Second)
	return j * 2
}

func main() {
	jobs := make(chan int)

	// buffered
	results := make(chan int, 9)

	// Start 3 worker Goroutines.
	for w := 1; w <= 3; w++ {
		go worker(jobs, results)
	}

	// Send 9 jobs.
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// Receive results.
	for a := 1; a <= 9; a++ {
		fmt.Println(<-results)
	}
}
