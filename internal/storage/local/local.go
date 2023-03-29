package local

import (
	"bufio"
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
	indexer *index.Indexer
	dir     string
	timeout time.Duration
}

func New(indexer *index.Indexer, dir string, timeout time.Duration) storage.Storage {
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
	if strings.HasSuffix(name, "/@v/list") {
		s := bufio.NewScanner(content)
		for s.Scan() {
			logging.Info(s.Text())
		}
	}
	return nil
}

func (l *Local) Clear() {
}
