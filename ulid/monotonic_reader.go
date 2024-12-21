package ulid

import (
	"io"
	"sync"
)

// MonotonicReader is an interface that should yield monotonically increasing entropy into the provided slice for all calls with the same ms parameter.
// If a MonotonicReader is provided to the New constructor, its MonotonicRead method will be used instead of Read.
type MonotonicReader interface {
	io.Reader
	MonotonicRead(ms uint64, p []byte) error
}

// LockedMonotonicReader wraps a MonotonicReader
// with a sync.Mutex for safe concurrent use.
type LockedMonotonicReader struct {
	mu sync.Mutex
	MonotonicReader
}

// MonotonicRead synchronizes calls to the wrapped MonotonicReader.
func (r *LockedMonotonicReader) MonotonicRead(ms uint64, p []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.MonotonicReader.MonotonicRead(ms, p)
}
