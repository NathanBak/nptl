package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NathanBak/nptl/src/nptl"
	"golang.org/x/text/language"
	"golang.org/x/text/message/pipeline"
)

type GoMessageHandler struct {
	Translator nptl.Translator
	NewLang    language.Tag
}

func (h *GoMessageHandler) Handle(ctx context.Context, source []byte) ([]byte, error) {

	ms := pipeline.Messages{}

	err := json.Unmarshal(source, &ms)
	if err != nil {
		return nil, err
	}

	if h.NewLang != language.Und {
		ms.Language = h.NewLang
	}

	for i, m := range ms.Messages {
		rs, err := nptl.FromString(m.Message.Msg)
		if err != nil {
			return nil, err
		}

		cache := SubstitutionCache{}

		for _, p := range m.Placeholders {
			rs, err = cache.SubstituteStrings(rs, fmt.Sprintf("{%s}", p.ID), p.String)
			if err != nil {
				return nil, err
			}
		}

		rs, err = h.Translator.Translate(ctx, rs)
		if err != nil {
			return nil, err
		}

		rs, err = cache.ReplaceAll(rs)
		if err != nil {
			return nil, err
		}

		ms.Messages[i].Translation = pipeline.Text{Msg: rs.String()}

	}

	return json.MarshalIndent(&ms, "", "   ")
}
