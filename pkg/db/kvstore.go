package db

import (
	"errors"

	badger "github.com/dgraph-io/badger/v4"
)

func NewKVStore(path string) (*KVStore, error) {
	opts := badger.DefaultOptions(path)

	db, err := badger.Open(opts)
	if err != nil {
		return nil, errors.Join(err, ErrFailedToOpenDB)
	}

	return &KVStore{db}, nil
}

func (k *KVStore) Get(key []byte) ([]byte, error) {
	var valCopy []byte
	err := k.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return errors.Join(err, ErrorKeyNotFound)
		}
		valCopy, err = item.ValueCopy(nil)
		return err
	})
	return valCopy, err
}

func (k *KVStore) Set(key, value []byte) error {
	return k.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

func (k *KVStore) Delete(key []byte) error {
	return k.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

func (k *KVStore) List() (map[string][]byte, error) {
	list := make(map[string][]byte)
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
			list[k] = v
		}
		return nil
	})
	return list, err
}

func (k *KVStore) Close() error {
	return k.db.Close()
}
