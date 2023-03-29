package cache

import (
	"context"
	"io"
	"time"
)

type Cache struct {
	dir     string
	timeout time.Duration
}

func New(dir string, timeout time.Duration) *Cache {
	return &Cache{
		dir:     dir,
		timeout: timeout,
	}
}

func (c *Cache) Get(ctx context.Context, name string) (io.ReadCloser, error) {
	return nil, nil
}

func (c *Cache) Set(ctx context.Context, name string, content io.ReadSeeker) error {
	return nil
}
