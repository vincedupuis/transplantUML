package model

import "fmt"

// Warnings collects the things a parser had to drop or an emitter could only
// approximate. They are reported to the user but do not stop the conversion.
type Warnings []string

// Addf appends a formatted warning.
func (w *Warnings) Addf(format string, args ...any) {
	*w = append(*w, fmt.Sprintf(format, args...))
}
