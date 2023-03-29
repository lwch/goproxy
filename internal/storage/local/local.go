package local

import (
	"context"
	"encoding/hex"
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
	if strings.HasSuffix(name, "/@v/list") ||
		strings.HasSuffix(name, "/@latest") {
		return l.Base.Get(ctx, name)
	}
	return nil, os.ErrNotExist
}

func (l *Local) Set(ctx context.Context, name string, content io.ReadSeeker) error {
	logging.Info("SET: %s", name)
	if strings.HasSuffix(name, "/@v/list") ||
		strings.HasSuffix(name, "/@latest") {
		return l.Base.Save(name, content)
	}
	data, err := io.ReadAll(content)
	if err != nil {
		return err
	}
	logging.Info("%s", hex.Dump(data))
	return nil
}

func (l *Local) Clear() {
}
