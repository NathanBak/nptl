package nptl

import (
	"context"
)

type Handler interface {
	Handle(ctx context.Context, source []byte) ([]byte, error)
}
