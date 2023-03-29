package index

import "github.com/dgraph-io/badger/v4"

type Indexer struct {
	db *badger.DB
}

func New(db *badger.DB) *Indexer {
	return &Indexer{db: db}
}
