package storage

import (
	"context"
	"goproxy/internal/index"
	"io"
)

type Storage interface {
	Get(ctx context.Context, name string) (io.ReadCloser, error)
	Put(ctx context.Context, name string, content io.ReadSeeker) error
	Clear()
}

type Base struct {
	indexer *index.Indexer
}

func New(indexer *index.Indexer) *Base {
	return &Base{indexer: indexer}
}

func (b *Base) Save(name string, content io.ReadSeeker) error {
	return b.indexer.Save(name, content)
}

func (b *Base) Get(_ context.Context, name string) (io.ReadCloser, error) {
	return b.indexer.Get(name)
}
