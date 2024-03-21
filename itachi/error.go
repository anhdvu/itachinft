package itachi

// Error represents system error.
// It can then be used with application layers
type Error struct {
	message string
}

// Error implements the error interface
func (e *Error) Error() string {
	return e.message
}
