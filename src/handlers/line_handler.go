package handlers

import (
	"bufio"
	"bytes"
	"context"
	"io"

	"github.com/NathanBak/nptl/src/nptl"
	"github.com/NathanBak/nptl/src/runes"
)

type LineHandler struct {
	Translator          nptl.Translator
	TranslateBlankLines bool
}

func (h *LineHandler) Handle(ctx context.Context, source []byte) ([]byte, error) {
	bytesReader := bytes.NewReader(source)
	br := bufio.NewReader(bytesReader)

	out := []byte{}

	for {
		line, nl, err := readLine(br)

		if err != nil && err != io.EOF {
			return out, err
		}

		if len(line) > 0 || h.TranslateBlankLines {
			rs, err2 := runes.FromBytes(line)
			if err2 != nil {
				return out, err2
			}

			rs, err2 = h.Translator.Translate(ctx, rs)
			if err2 != nil {
				return out, err2
			}

			buf, err2 := rs.ToUTF8Bytes()
			if err2 != nil {
				return out, err2
			}

			out = append(out, buf...)
		}
		out = append(out, nl...)

		if err == io.EOF {
			break
		}
	}

	return out, nil
}
