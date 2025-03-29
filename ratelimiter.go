package golimit

import (
	"sync"
	"time"
)

type RateLimiter struct {
	tokens     int
	capacity   int
	burst      int
	refillRate time.Duration
	lastRefill time.Time
	mu         sync.Mutex
}

func NewRateLimiter(capacity, burst int, refillRate time.Duration) *RateLimiter {
	if burst > capacity {
		burst = capacity
	}

	if refillRate <= 0 {
		panic("refillRate must be greater than zero")
	}

	return &RateLimiter{
		tokens:     burst,
		capacity:   capacity,
		burst:      burst,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

func (tb *RateLimiter) refill() {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)

	newTokens := int(elapsed / tb.refillRate)

	if newTokens > 0 {
		tb.tokens += newTokens

		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}

		tb.lastRefill = now
	}

}

func (tb *RateLimiter) Allow() bool {
	tb.refill()
	tb.mu.Lock()
	defer tb.mu.Unlock()

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}
