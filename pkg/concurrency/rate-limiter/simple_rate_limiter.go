package rate_limiter

import (
	"fmt"
	"time"
)

// NOTE : This is not a impl. of a distributed rate limiter.

func TickerBasedRateLimiter() {

	// buffered channel - which is a non-blocking channel
	requestQueueChan := make(chan int, 5)

	for i := 0; i < 5; i++ {
		// RS, sending integer to requests channel
		requestQueueChan <- i
	}
	close(requestQueueChan)

	// Below impl. can leak memory - so pls be careful to no use for long running tasks.
	limiter := time.Tick(time.Second)

	// closing the channel - will not delete the data which
	// is already sent.
	// We are ranging over the request chan.
	for req := range requestQueueChan {
		// limiter is obstructing it to call the function at a constant rate.
		<-limiter
		fmt.Println("requests", req, time.Now())
	}
}

func BurstyRateLimiter() {

	// -------- This section is maintaining a buffer for  burst
	//and then periodic sending to channel.

	// buffered channel with 3 limit
	// Allows burst of request.
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		// Filling the queue (of 3) with time data.
		burstyLimiter <- time.Now()
	}

	// sending the time data continuously - in periodic manner in a bg goroutine.
	go func() {
		for t := range time.Tick(time.Second) {
			burstyLimiter <- t
		}
	}()

	// -------- End of Section ---------

	// -------- This section is maintaining requests which are hit continuously.
	// 3 are allowed
	// bursty Requests
	burstyRequests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
