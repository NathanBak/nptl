package handlers

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/NathanBak/nptl/src/translators"
)

func TestKeyValueHandler01(t *testing.T) {
	ctx := context.Background()
	srcFile := "/home/bakchoy/vscode/src/github.com/NathanBak/nptl/test3.txt"
	tarFile := "/home/bakchoy/vscode/src/github.com/NathanBak/nptl/output.txt"

	buf, err := ioutil.ReadFile(srcFile)
	if err != nil {
		t.Error(err)
	}

	handler := KeyValueHandler{Translator: &translators.AccentTranslator{}}
	// handler.TranslateLeadingWhitespace = true

	out, err := handler.Handle(ctx, buf)
	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile(tarFile, out, 0644)
	if err != nil {
		t.Error(err)
	}
}

func bytesToRunes(buf []byte) ([]rune, error) {
	runes := []rune{}

	for _, r := range string(buf) {
		runes = append(runes, r)
	}
	return runes, nil
}
