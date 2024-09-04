package rate_limiter

import (
	"fmt"
	"go.uber.org/ratelimit"
	"time"
)

func UberRateLimiter() {
	rl := ratelimit.New(100)
	rl.Take()
	time.Sleep(time.Millisecond * 45)

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		if i > 0 {
			fmt.Println(i, now.Sub(prev).Round(time.Millisecond*2))
		}
		prev = now
	}
}
