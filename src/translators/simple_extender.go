package translators

import (
	"context"

	"github.com/NathanBak/nptl/src/nptl"
)

// SimpleExtender sticks the prefix on the beginning, the suffix on the end, and then inserts the
// extender rune until the new target is at least 20% longer than the original source Runes.
type SimpleExtender struct {
	Prefix          nptl.Runes
	Suffix          nptl.Runes
	Extender        rune
	ExtendFromFront bool
}

// Extend implements the Extender.Extend() method.
func (e *SimpleExtender) Extend(ctx context.Context, source, target []rune) ([]rune, error) {

	numberToAdd := 0
	if e.Extender != 0 {
		targetLength := int((float32(len(source)) * 1.2))
		numberToAdd = targetLength - len(source) - len(e.Prefix) - len(e.Suffix)
		if numberToAdd < 1 {
			numberToAdd = 1
		}
	}

	response := []rune{}
	response = append(response, e.Prefix...)

	if e.ExtendFromFront {
		for i := 0; i < numberToAdd; i++ {
			response = append(response, e.Extender)
		}
	}

	response = append(response, target...)

	if !e.ExtendFromFront {
		for i := 0; i < numberToAdd; i++ {
			response = append(response, e.Extender)
		}
	}

	response = append(response, e.Suffix...)

	return response, nil
}
