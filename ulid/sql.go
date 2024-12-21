package ulid

import (
	"database/sql/driver"

	"github.com/fajarnugraha37/goid/errors"
)

// Scan implements the sql.Scanner interface. It supports scanning a string or byte slice.
func (id *ULID) Scan(src interface{}) error {
	switch x := src.(type) {
	case nil:
		return nil
	case string:
		return id.UnmarshalText([]byte(x))
	case []byte:
		return id.UnmarshalBinary(x)
	}

	return errors.ErrUlidScanValue
}

// Value implements the sql/driver.Valuer interface, returning the ULID as a
// slice of bytes, by invoking MarshalBinary. If your use case requires a string
// representation instead, you can create a wrapper type that calls String() instead.
func (id *ULID) Value() (driver.Value, error) {
	return id.MarshalBinary()
}
