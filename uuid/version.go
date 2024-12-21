package uuid

import "fmt"

// A Version represents a UUID's version.
type Version byte

const (
	_ byte = iota
	V1
	V2
	V3
	V4
	V5
)

func (v Version) String() string {
	if v > 15 {
		return fmt.Sprintf("BAD_VERSION_%d", v)
	}
	return fmt.Sprintf("VERSION_%d", v)
}

// Version returns the version of uuid.
func (uuid UUID) Version() Version {
	return Version(uuid[6] >> 4)
}

// SetVersion sets version bits.
func (uuid *UUID) SetVersion(v byte) {
	uuid[6] = (uuid[6] & 0x0f) | (v << 4)
}
