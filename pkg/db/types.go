package db

import (
	"errors"

	badger "github.com/dgraph-io/badger/v4"
)

const defaultThreshold = 1024 * 1024 * 5 // 5MB

var ErrFailedToOpenDB = errors.New("failed to open db for kv store")

type KVStore struct {
	db *badger.DB
}

func NewDB(path string) (*KVStore, error) {
	opts := badger.DefaultOptions(path)
	opts = opts.WithValueThreshold(5242880)

	db, err := badger.Open(opts)
	if err != nil {
		return nil, errors.Join(err, ErrFailedToOpenDB)
	}
	defer db.Close()

	return &KVStore{db: db}, nil
}

func (k *KVStore) Get(key string) (string, error) {
	var valCopy []byte
	err := k.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		valCopy, err = item.ValueCopy(nil)
		return err
	})
	return string(valCopy), err
}

func (k *KVStore) Set(key, value string) error {
	return k.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
}

func (k *KVStore) Delete(key string) error {
	return k.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

func (k *KVStore) List() (map[string]string, error) {
	list := make(map[string]string)
	err := k.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := string(item.Key())
			var v []byte
			item.Value(func(val []byte) error {
				v = append([]byte{}, val...)
				return nil
			})
			list[k] = string(v)
		}
		return nil
	})
	return list, err
}

func (k *KVStore) Close() error {
	return k.db.Close()
}
