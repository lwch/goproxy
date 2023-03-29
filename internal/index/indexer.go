package index

import (
	"io"
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
