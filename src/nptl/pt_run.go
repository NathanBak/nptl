package nptl

import (
	"context"
	"io/ioutil"
	"os"
)

type PTCombination struct {
	SourceFile        string
	TargetFile        string
	Handler           Handler
	TargetPermissions os.FileMode
}

func (r PTCombination) PseudoTranslate(ctx context.Context) error {

	buf, err := ioutil.ReadFile(r.SourceFile)
	if err != nil {
		return err
	}

	out, err := r.Handler.Handle(ctx, buf)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.TargetFile, out, 0644)
	if err != nil {
		return err
	}

	return nil
}
