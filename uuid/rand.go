package uuid

import (
	"crypto/rand"
	"io"
	"sync"
)

const randPoolSize = 16 * 16

var (
	// Zero is special form of UUID that is specified to have all
	// 128 bits set to zero.
	rander      = rand.Reader
	poolEnabled = false
	poolMu      sync.Mutex
	poolPos     = randPoolSize     // protected with poolMu
	pool        [randPoolSize]byte // protected with poolMu
)

// SetRand sets the random number generator to r, which implements io.Reader.
// If r.Read returns an error when the package requests random data then
// a panic will be issued.
//
// Calling SetRand with nil sets the random number generator to the default
// generator.
func SetRand(r io.Reader) {
	if r == nil {
		rander = rand.Reader
		return
	}
	rander = r
}

// EnableRandPool enables internal randomness pool used for Random
// (Version 4) UUID generation. The pool contains random bytes read from
// the random number generator on demand in batches. Enabling the pool
// may improve the UUID generation throughput significantly.
//
// Since the pool is stored on the Go heap, this feature may be a bad fit
// for security sensitive applications.
//
// Both EnableRandPool and DisableRandPool are not thread-safe and should
// only be called when there is no possibility that New or any other
// UUID Version 4 generation function will be called concurrently.
func EnableRandPool() {
	poolEnabled = true
}

// DisableRandPool disables the randomness pool if it was previously
// enabled with EnableRandPool.
//
// Both EnableRandPool and DisableRandPool are not thread-safe and should
// only be called when there is no possibility that New or any other
// UUID Version 4 generation function will be called concurrently.
func DisableRandPool() {
	poolEnabled = false
	defer poolMu.Unlock()
	poolMu.Lock()
	poolPos = randPoolSize
}
