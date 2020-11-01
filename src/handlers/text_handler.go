package handlers

import (
	"context"

	"github.com/NathanBak/nptl/src/nptl"
	"github.com/NathanBak/nptl/src/runes"
)

type TextHandler struct {
	Translator nptl.Translator
}

func (h *TextHandler) Handle(ctx context.Context, source []byte) ([]byte, error) {

	rs, err := runes.FromBytes(source)
	if err != nil {
		return nil, err
	}

	rs, err = h.Translator.Translate(ctx, rs)
	if err != nil {
		return nil, err
	}

	return rs.ToUTF8Bytes()
}
