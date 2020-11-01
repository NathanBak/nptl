package handlers

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"unicode"

	"github.com/NathanBak/nptl/src/nptl"
	"github.com/NathanBak/nptl/src/runes"
)

type KeyValueHandler struct {
	Translator                 nptl.Translator
	TranslateLeadingWhitespace bool
}

func (h *KeyValueHandler) Handle(ctx context.Context, source []byte) ([]byte, error) {
	bytesReader := bytes.NewReader(source)
	br := bufio.NewReader(bytesReader)

	out := []byte{}

	for {
		line, nl, err := readLine(br)

		if err != nil && err != io.EOF {
			return out, err
		}

		rs, err2 := runes.FromBytes(line)
		if err2 != nil {
			return out, err2
		}

		k, v, err2 := h.split(rs)
		if err2 != nil {
			return out, err2
		}

		if len(v) > 0 {
			v, err2 = h.Translator.Translate(ctx, v)
			if err2 != nil {
				return out, err2
			}
		}

		rs = k
		rs = append(rs, v...)

		buf, err2 := rs.ToUTF8Bytes()
		if err2 != nil {
			return out, err2
		}

		out = append(out, buf...)
		out = append(out, nl...)

		if err == io.EOF {
			break
		}
	}

	return out, nil
}

func (h *KeyValueHandler) split(rs runes.Runes) (runes.Runes, runes.Runes, error) {
	key := runes.Runes{}
	val := runes.Runes{}

	foundEquals := false
	keyDone := false

	for _, r := range rs {
		if !foundEquals {
			key = append(key, r)
		}

		if foundEquals && !keyDone {
			if unicode.IsSpace(r) {
				key = append(key, r)
				continue
			}
			keyDone = true
		}

		if !foundEquals && r == '=' {
			foundEquals = true
			keyDone = h.TranslateLeadingWhitespace
			continue
		}

		if keyDone {
			val = append(val, r)
		}
	}

	return key, val, nil
}
