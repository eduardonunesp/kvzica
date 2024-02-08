package db

import (
	"errors"

	badger "github.com/dgraph-io/badger/v4"
)

var (
	ErrFailedToOpenDB = errors.New("failed to open db for kv store")
	ErrorKeyNotFound  = errors.New("key not found")
)

type KVStore struct {
	db *badger.DB
}
