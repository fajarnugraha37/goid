package uuid

// Size of a UUID in bytes.
const Size = 16

// A UUID is a 128 bit (16 byte) Universal Unique IDentifier as defined in RFC 9562.
type UUID [Size]byte

// String returns the string form of uuid, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// , or "" if uuid is invalid.
func (uuid UUID) String() string {
	var buf [36]byte
	encodeHex(buf[:], uuid)
	return string(buf[:])
}

// URN returns the RFC 2141 URN form of uuid,
// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx, or "" if uuid is invalid.
func (uuid UUID) URN() string {
	var buf [36 + 9]byte
	copy(buf[:], "urn:uuid:")
	encodeHex(buf[9:], uuid)
	return string(buf[:])
}
