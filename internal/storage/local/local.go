package local

import (
	"context"
	"goproxy/internal/storage"
	"io"
	"os"
	"time"

	"github.com/lwch/logging"
)

type Local struct {
	dir     string
	timeout time.Duration
}

func New(dir string, timeout time.Duration) storage.Storage {
	return &Local{
		dir:     dir,
		timeout: timeout,
	}
}

func (l *Local) Get(ctx context.Context, name string) (io.ReadCloser, error) {
	logging.Info("GET: %s", name)
	return nil, os.ErrNotExist
}

func (l *Local) Set(ctx context.Context, name string, content io.ReadSeeker) error {
	logging.Info("SET: %s", name)
	return nil
}

func (l *Local) Clear() {
}
