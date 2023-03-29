package cache

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/lwch/logging"
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
	logging.Info("GET: %s", name)
	return nil, os.ErrNotExist
}

func (c *Cache) Set(ctx context.Context, name string, content io.ReadSeeker) error {
	logging.Info("SET: %s", name)
	return nil
}
