package uuid

import "fmt"

// A Variant represents a UUID's variant.
type Variant byte

// Constants returned by Variant.
const (
	Invalid   = Variant(iota) // Invalid UUID
	RFC4122                   // The variant specified in RFC9562(obsoletes RFC4122).
	Reserved                  // Reserved, NCS backward compatibility.
	Microsoft                 // Reserved, Microsoft Corporation backward compatibility.
	Future                    // Reserved for future definition.
)

// RFC9562 added V6 and V7 of UUID, but did not change specification of V1 and V4
// implemented in this module. To avoid creating new major module version,
// we still use RFC4122 for constant name.
const Standard = RFC4122

func (v Variant) String() string {
	switch v {
	case RFC4122:
		return "RFC4122"
	case Reserved:
		return "Reserved"
	case Microsoft:
		return "Microsoft"
	case Future:
		return "Future"
	case Invalid:
		return "Invalid"
	}
	return fmt.Sprintf("BadVariant%d", int(v))
}

// Variant returns the variant encoded in uuid.
func (uuid UUID) Variant() Variant {
	switch {
	case (uuid[8] & 0xc0) == 0x80:
		return RFC4122
	case (uuid[8] & 0xe0) == 0xc0:
		return Microsoft
	case (uuid[8] & 0xe0) == 0xe0:
		return Future
	default:
		return Reserved
	}
}

// SetVariant sets variant bits.
func (u *UUID) SetVariant(v Variant) {
	switch v {
	case Reserved:
		u[8] = (u[8]&(0xff>>1) | (0x00 << 7))
	case RFC4122:
		u[8] = (u[8]&(0xff>>2) | (0x02 << 6))
	case Microsoft:
		u[8] = (u[8]&(0xff>>3) | (0x06 << 5))
	case Future:
		fallthrough
	default:
		u[8] = (u[8]&(0xff>>3) | (0x07 << 5))
	}
}
