package nptl

import (
	"context"
)

// A Translator is able to apply a pseudo translation algorithm to Runes.
type Translator interface {
	// Translate accepts Runes and returns a pseudo translated variant.
	Translate(context.Context, Runes) (Runes, error)
}
