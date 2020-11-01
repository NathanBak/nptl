package translators

import (
	"context"
	"sync"

	"github.com/NathanBak/nptl/src/nptl"
)

// The AccentTranslator simulates translation into Western European languages.
type AccentTranslator struct {
	extender  Extender
	setupOnce sync.Once
}

// Translate implements the nptl.Translator.Translate() method.
func (t *AccentTranslator) Translate(ctx context.Context, source nptl.Runes) (nptl.Runes, error) {
	t.setupOnce.Do(func() {
		t.extender = &SimpleExtender{
			Prefix:   nptl.Runes{'['},
			Suffix:   nptl.Runes{'€', ']'},
			Extender: '~',
		}
	})

	response := []rune{}

	for _, r := range source {
		v, ok := replacements[r]
		if ok {
			response = append(response, v)
		} else {
			response = append(response, r)
		}
	}

	return t.extender.Extend(ctx, source, response)
}

var replacements = map[rune]rune{
	'a': 'á',
	'c': 'ç',
	'e': 'è',
	'i': 'î',
	'n': 'ñ',
	'o': 'ö',
	'u': 'û',
	'y': 'ý',

	'A': 'Å',
	'C': 'Ç',
	'E': 'È',
	'I': 'Ï',
	'D': 'Ð',
	'N': 'Ñ',
	'O': 'Ø',
	'U': 'Û',
	'Y': 'Ý',
}
