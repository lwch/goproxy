package storage

import (
	"context"
	"io"
)

type Storage interface {
	Get(ctx context.Context, name string) (io.ReadCloser, error)
	Set(ctx context.Context, name string, content io.ReadSeeker) error
	Clear()
}
