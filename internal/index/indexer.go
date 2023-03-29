package index

import (
	"bytes"
	"io"
	"os"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/lwch/logging"
)

type Indexer struct {
	db      *badger.DB
	timeout time.Duration
}

func New(db *badger.DB, timeout time.Duration) *Indexer {
	return &Indexer{db: db, timeout: timeout}
}

func (idx *Indexer) Save(name string, r io.Reader) error {
	return idx.db.Update(func(txn *badger.Txn) error {
		data, err := io.ReadAll(r)
		if err != nil {
			logging.Info("read data: %v", err)
			return nil
		}
		e := badger.NewEntry([]byte(name), data)
		return txn.SetEntry(e)
	})
}

func (idx *Indexer) Get(name string) (io.ReadCloser, error) {
	var data []byte
	err := idx.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(name))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return os.ErrNotExist
			}
			return err
		}
		return item.Value(func(val []byte) error {
			data = make([]byte, len(val))
			copy(data, val)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return io.NopCloser(bytes.NewReader(data)), nil
}
