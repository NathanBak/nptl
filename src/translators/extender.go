package translators

import "context"

// An Extender is used to expand pseudo translated Runes to simulate more verbose languages.  It
// may also add markers at the beginning and end to help identify concatenation.
type Extender interface {
	// Extend modifies Runes to be longer.
	Extend(ctx context.Context, source, target []rune) ([]rune, error)
}
