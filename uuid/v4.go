package uuid

import "io"

// NewV4 creates a new random UUID or panics.
func NewV4() UUID {
	return Must(NewV4Random())
}

// NewV4String creates a new random UUID and returns it as a string or panics.
func NewV4String() string {
	return Must(NewV4Random()).String()
}

// NewV4Random returns a Random (Version 4) UUID.
//
// The strength of the UUIDs is based on the strength of the crypto/rand
// package.
//
// Uses the randomness pool if it was enabled with EnableRandPool.
//
// A note about uniqueness derived from the UUID Wikipedia entry:
//
//  Randomly generated UUIDs have 122 random bits.  One's annual risk of being
//  hit by a meteorite is estimated to be one chance in 17 billion, that
//  means the probability is about 0.00000000006 (6 × 10−11),
//  equivalent to the odds of creating a few tens of trillions of UUIDs in a
//  year and having one duplicate.
func NewV4Random() (UUID, error) {
	if !poolEnabled {
		return NewV4RandomFromReader(rander)
	}
	return newV4RandomFromPool()
}

// NewV4RandomFromReader returns a UUID based on bytes read from a given io.Reader.
func NewV4RandomFromReader(r io.Reader) (UUID, error) {
	var uuid UUID
	_, err := io.ReadFull(r, uuid[:])
	if err != nil {
		return Nil, err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid, nil
}

func newV4RandomFromPool() (UUID, error) {
	var uuid UUID
	poolMu.Lock()
	if poolPos == randPoolSize {
		_, err := io.ReadFull(rander, pool[:])
		if err != nil {
			poolMu.Unlock()
			return Nil, err
		}
		poolPos = 0
	}
	copy(uuid[:], pool[poolPos:(poolPos+16)])
	poolPos += 16
	poolMu.Unlock()

	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid, nil
}
