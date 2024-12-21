package ulid

import (
	"bytes"

	"github.com/fajarnugraha37/goid/errors"
)

const (
	// Encoding is the base 32 encoding alphabet used in ULID strings.
	Encoding = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
	// A ULID consists of 26 characters, which includes:
	// - A 48-bit timestamp (milliseconds since Unix epoch).
	// - A 80-bit random component.
	EncodedSize = 26
)

/*
A ULID is a 16 byte Universally Unique Lexicographically Sortable Identifier

	The components are encoded as 16 octets.
	Each component is encoded with the MSB first (network byte order).

	0                   1                   2                   3
	0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	|                      32_bit_uint_time_high                    |
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	|     16_bit_uint_time_low      |       16_bit_uint_random      |
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	|                       32_bit_uint_random                      |
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	|                       32_bit_uint_random                      |
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
*/
type ULID [16]byte

var (
	// Zero is a zero-value ULID.
	Zero ULID
	// maxTime is the maximum Unix time in milliseconds that can be represented in a ULID.
	maxTime = (&ULID{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}).Time()
	// MaxTime returns the maximum Unix time in milliseconds that can be encoded in a ULID.
	MaxTime = func() uint64 {
		return maxTime
	}
)

// Bytes returns bytes slice representation of ULID.
func (id *ULID) Bytes() []byte {
	return id[:]
}

// String returns a lexicographically sortable string encoded ULID (26 characters, non-standard base 32) e.g. 01AN4Z07BY79KA1307SR9X4MV3.
// Format: tttttttttteeeeeeeeeeeeeeee where t is time and e is entropy.
func (id *ULID) String() string {
	ulid := make([]byte, EncodedSize)
	_ = id.MarshalTextTo(ulid)
	return string(ulid)
}

// IsZero returns true if the ULID is a zero-value ULID, i.e. ulid.Zero.
func (id *ULID) IsZero() bool {
	return id.Compare(Zero) == 0
}

// Entropy returns the entropy from the ULID.
func (id *ULID) Entropy() []byte {
	e := make([]byte, 10)
	copy(e, id[6:])
	return e
}

// SetEntropy sets the ULID entropy to the passed byte slice.
// ErrDataSize is returned if len(e) != 10.
func (id *ULID) SetEntropy(e []byte) error {
	if len(e) != 10 {
		return errors.ErrUlidDataSize
	}

	copy((*id)[6:], e)
	return nil
}

// Compare returns an integer comparing id and other lexicographically.
// The result will be 0 if id==other, -1 if id < other, and +1 if id > other.
func (id *ULID) Compare(other ULID) int {
	return bytes.Compare(id[:], other[:])
}
