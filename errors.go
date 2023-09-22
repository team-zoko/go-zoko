package gozoko

import "fmt"

type ArgError struct {
	arg    string
	reason string
}

var _ error = &ArgError{}

// NewArgError creates an InputError.
func NewArgError(arg, reason string) *ArgError {
	return &ArgError{
		arg:    arg,
		reason: reason,
	}
}

func (e *ArgError) Error() string {
	return fmt.Sprintf("%s is invalid because %s", e.arg, e.reason)
}
