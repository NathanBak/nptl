package nptl

import "context"

type Translator interface {
	Translate(context.Context, []rune) ([]rune, error)
}
