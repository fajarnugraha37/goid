package errors

import (
	e "errors"
	"fmt"
)

var (
	ErrInvalidUUIDFormat      = e.New("[UUID] invalid UUID format")
	ErrInvalidBracketedFormat = e.New("[UUID] invalid bracketed UUID format")
	ErrInvalidURNPrefix       = URNPrefixError{}
	ErrInvalidLength          = InvalidLengthError{}
)

type URNPrefixError struct {
	Prefix string
}

func (e URNPrefixError) Error() string {
	return fmt.Sprintf("[UUID] invalid urn prefix: %q", e.Prefix)
}

func (e URNPrefixError) Is(target error) bool {
	_, ok := target.(URNPrefixError)
	return ok
}

type InvalidLengthError struct {
	Len int
}

func (err InvalidLengthError) Error() string {
	return fmt.Sprintf("[UUID] invalid UUID length: %d", err.Len)
}

func (e InvalidLengthError) Is(target error) bool {
	_, ok := target.(InvalidLengthError)
	return ok
}
