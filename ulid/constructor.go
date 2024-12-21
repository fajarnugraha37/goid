package ulid

import (
	"io"
	"time"
)

// New returns a ULID with the given Unix milliseconds timestamp and an optional entropy source.
// Use the Timestamp function to convert a time.Time to Unix milliseconds.
//
// ErrBigTime is returned when passing a timestamp bigger than MaxTime.
// Reading from the entropy source may also return an error.
//
// Safety for concurrent use is only dependent on the safety of the entropy source.
func New(ms uint64, entropy io.Reader) (*ULID, error) {
	var (
		id  ULID
		err error
	)
	if err = id.SetTime(ms); err != nil {
		return &id, err
	}

	switch e := entropy.(type) {
	case nil:
		return &id, err
	case MonotonicReader:
		err = e.MonotonicRead(ms, id[6:])
	default:
		_, err = io.ReadFull(e, id[6:])
	}

	return &id, err
}

// MustNew is a convenience function equivalent to New that panics on failure instead of returning an error.
func MustNew(ms uint64, entropy io.Reader) *ULID {
	id, er := New(ms, entropy)
	if er != nil {
		panic(er)
	}
	return id
}

// MustNewDefault is a convenience function equivalent to MustNew with DefaultEntropy as the entropy.
// It may panic if the given time.Time is too large or too small.
func MustNewDefault(t time.Time) *ULID {
	return MustNew(Timestamp(t), defaultEntropy)
}

// Make returns a ULID with the current time in Unix milliseconds and monotonically increasing entropy for the same millisecond.
// It is safe for concurrent use, leveraging a sync.Pool underneath for minimal contention.
func Make() *ULID {
	// NOTE: MustNew can't panic since DefaultEntropy never returns an error.
	return MustNew(Now(), defaultEntropy)
}

// Parse parses an encoded ULID, returning an error in case of failure.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded  ULID's length.
// Invalid encodings produce undefined ULIDs. For a version that returns an error instead, see ParseStrict.
func Parse(ulid string) (*ULID, error) {
	var id ULID
	return &id, parse([]byte(ulid), false, &id)
}

// ParseStrict parses an encoded ULID, returning an error in case of failure.
//
// It is like Parse, but additionally validates that the parsed ULID consists only of valid base32 characters. It is slightly slower than Parse.
//
// ErrDataSize is returned if the len(ulid) is different from an encoded ULID's length.
// Invalid encodings return ErrInvalidCharacters.
func ParseStrict(ulid string) (*ULID, error) {
	var id ULID
	return &id, parse([]byte(ulid), true, &id)
}

// MustParse is a convenience function equivalent to Parse that panics on failure instead of returning an error.
func MustParse(ulid string) *ULID {
	id, er := Parse(ulid)
	if er != nil {
		panic(er)
	}
	return id
}

// MustParseStrict is a convenience function equivalent to ParseStrict that panics on failure instead of returning an error.
func MustParseStrict(ulid string) *ULID {
	id, er := ParseStrict(ulid)
	if er != nil {
		panic(er)
	}
	return id
}
