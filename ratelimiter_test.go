package golimit

import (
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {

	tb := NewRateLimiter(10, 5, time.Second)

	for i := 0; i < 5; i++ {
		if !tb.Allow() {
			t.Error("Expected request to be allowed but it was denied")
		}
	}

}

func TestTokenBucketRefill(t *testing.T) {

	tb := NewRateLimiter(10, 5, time.Second)

	for i := 0; i < 5; i++ {
		if !tb.Allow() {
			t.Errorf("Expected %d requests to be allowed but it was denied", i+1)
		}
	}

	if tb.Allow() {
		t.Errorf("Expected requests to be denied but it was allowed. Number of tokens left: %d", tb.tokens)
	}

	time.Sleep(3 * time.Second)

	for i := 0; i < 3; i++ {
		if !tb.Allow() {
			t.Errorf("Expected %d requests to be allowed but it was denied", i+1)
		}
	}

	if tb.Allow() {
		t.Errorf("Expected request to be denied but it was allowed. Number of tokens left: %d", tb.tokens)
	}
}
