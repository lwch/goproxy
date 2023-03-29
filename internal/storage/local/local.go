package local

import (
	"context"
	"goproxy/internal/index"
	"goproxy/internal/storage"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
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
	if strings.HasSuffix(name, ".zip") {
		tmp := strings.SplitN(name, "/@v/", 2)
		dir := filepath.Join(l.dir, tmp[0], tmp[1])
		f, err := os.Open(dir)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		fi, err := f.Stat()
		if err != nil {
			return nil, err
		}
		return &struct {
			*os.File
			os.FileInfo
		}{f, fi}, nil
	}
	return l.Base.Get(ctx, name)
}

func (l *Local) Set(ctx context.Context, name string, content io.ReadSeeker) error {
	logging.Info("SET: %s", name)
	if strings.HasSuffix(name, ".zip") {
		tmp := strings.SplitN(name, "/@v/", 2)
		dir := filepath.Join(l.dir, tmp[0], tmp[1])
		err := os.MkdirAll(filepath.Dir(dir), 0755)
		if err != nil {
			return err
		}
		f, err := ioutil.TempFile(filepath.Dir(dir), "tmp")
		if err != nil {
			return err
		}
		defer os.Remove(f.Name())
		_, err = io.Copy(f, content)
		if err != nil {
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}
		return os.Rename(f.Name(), dir)
	}
	return l.Base.Save(name, content)
}

func (l *Local) Clear() {
	filepath.Walk(l.dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if time.Since(info.ModTime()) > l.timeout {
			os.Remove(path)
		}
		return nil
	})
}
