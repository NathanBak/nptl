package handlers

import (
	"context"

	"github.com/NathanBak/nptl/src/nptl"
)

type TextHandler struct {
	Translator nptl.Translator
}

func (h *TextHandler) Handle(ctx context.Context, source []byte) ([]byte, error) {

	rs, err := nptl.FromBytes(source)
	if err != nil {
		return nil, err
	}

	rs, err = h.Translator.Translate(ctx, rs)
	if err != nil {
		return nil, err
	}

	return rs.ToUTF8Bytes()
}
