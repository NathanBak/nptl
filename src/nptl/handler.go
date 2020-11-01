package nptl

import (
	"context"
)

// A Handler is able to pseudo translate the translatble content for a specific format.
type Handler interface {
	// Handle takes in the source bytes (usually from a file) and returns modified target bytes
	// where the translatable content has been pseudo translated.
	Handle(ctx context.Context, source []byte) ([]byte, error)
}
