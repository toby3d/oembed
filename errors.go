package oembed

import (
	"fmt"

	"golang.org/x/xerrors"
)

// Error represent a complex error
type Error struct {
	Message string
	URL     string
	Details xerrors.Frame
}

// Error returns a string formatted error
func (e Error) Error() string {
	return fmt.Sprint(e)
}

// Format implements fmt.Formatter method
func (e Error) Format(f fmt.State, c rune) {
	xerrors.FormatError(e, f, c)
}

// FormatError implements xerrors.Formatter method
func (e Error) FormatError(p xerrors.Printer) error {
	p.Printf("ERROR: %d [url:%s]", e.Message, e.URL)
	if p.Detail() {
		e.Details.Format(p)
	}
	return nil
}
