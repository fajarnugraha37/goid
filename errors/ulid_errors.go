package errors

import (
	e "errors"
)

var (
	// ErrUlidDataSize is returned when parsing or unmarshaling ULIDs with the wrong data size.
	ErrUlidDataSize = e.New("[ULID] bad data size when unmarshaling")
	// ErrUlidInvalidCharacters is returned when parsing or unmarshaling ULIDs with invalid Base32 encodings.
	ErrUlidInvalidCharacters = e.New("[ULID] bad data characters when unmarshaling")
	// ErrUlidBufferSize is returned when marshalling ULIDs to a buffer of insufficient size.
	ErrUlidBufferSize = e.New("[ULID] bad buffer size when marshaling")
	// ErrUlidBigTime is returned when constructing a ULID with a time that is larger than MaxTime.
	ErrUlidBigTime = e.New("[ULID] time too big")
	// ErrUlidOverflow is returned when unmarshaling a ULID whose first character is larger than 7, thereby exceeding the valid bit depth of 128.
	ErrUlidOverflow = e.New("[ULID] overflow when unmarshaling")
	// ErrUlidMonotonicOverflow is returned by a Monotonic entropy source when incrementing the previous ULID's entropy bytes would result in overflow.
	ErrUlidMonotonicOverflow = e.New("[ULID] monotonic entropy overflow")
	// ErrUlidScanValue is returned when the value passed to scan cannot be unmarshaled into the ULID.
	ErrUlidScanValue = e.New("[ULID] source value must be a string or byte slice")
)
