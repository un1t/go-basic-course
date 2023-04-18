package foo

import (
	"sync"
	"sync/atomic"
)

var mu sync.Mutex

func IncrementMutext(x *uint64) {
	mu.Lock()
	*x += 1
	mu.Unlock()
}

func IncrementAtomic(x *uint64) {
	atomic.AddUint64(x, 1)
}
