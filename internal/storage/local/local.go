package local

import (
	"context"
	"goproxy/internal/index"
	"goproxy/internal/storage"
	"io"
	"os"
	"strings"
	"time"

	"github.com/lwch/logging"
)

type Local struct {
	*storage.Base
	dir     string
	timeout time.Duration
}

func New(indexer *index.Indexer, dir string, timeout time.Duration) storage.Storage {
	return &Local{
		Base:    storage.New(indexer),
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
	if strings.HasSuffix(name, "/@v/list") {
		return l.Base.Save(strings.TrimSuffix(name, "/@v/list"), content)
	}
	return nil
}

func (l *Local) Clear() {
}
