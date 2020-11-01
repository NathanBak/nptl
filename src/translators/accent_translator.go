package translators

import (
	"context"
	"sync"
)

type AccentTranslator struct {
	extender  Extender
	setupOnce sync.Once
}

func (t *AccentTranslator) Translate(ctx context.Context, source []rune) ([]rune, error) {
	t.setupOnce.Do(func() {
		t.extender = &SimpleExtender{}
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

	return t.extender.Extend(ctx, source, response, []rune{'['}, []rune{'€', ']'}, '~')
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
