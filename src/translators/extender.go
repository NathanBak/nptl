package translators

import "context"

type Extender interface {
	Extend(ctx context.Context, source, target, prefix, suffix []rune, extenderRune rune) ([]rune, error)
}
