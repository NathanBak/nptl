package main

import (
	"context"
	"log"
	"runtime"
	"strings"

	"github.com/NathanBak/nptl/src/handlers"
	"github.com/NathanBak/nptl/src/nptl"
	"github.com/NathanBak/nptl/src/translators"
)

func main() {

	translator := &translators.AccentTranslator{}
	handler := &handlers.KeyValueHandler{Translator: translator}

	ptc := nptl.PTCombination{
		SourceFile: getFullPath("example_in.txt"),
		TargetFile: getFullPath("example_out.txt"),
		Handler:    handler,
	}

	err := ptc.PseudoTranslate(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func getFullPath(in string) string {
	_, filename, _, _ := runtime.Caller(0)
	lastSlash := strings.LastIndex(filename, "/")

	return filename[:lastSlash+1] + in
}
