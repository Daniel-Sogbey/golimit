# GoLimit

A lightweight and efficient rate limiter, implemented using the token bucket algorithm

## Installation
`` go get github.com/Daniel-Sogbey/golimit
``

## Usage

```go
import (
"fmt",
"time",
"github.com/Daniel-Sogbey/golimit"
)

func main(){
    tb := NewRateLimiter(10, 5, 1)
	
    for i:=0; i< 5; i++ {
        if tb.Allow() {
            fmt.Println("request allowed")
        } else {
            fmt.Println("request denied")
        }   
        time.Sleep(200 * time.Millisecond)
    }
}
```

