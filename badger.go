package dictionary

import (
	"io"

	badger "github.com/dgraph-io/badger/v3"
)

type BadgerDB struct {
	data string
	db   *badger.DB
}

func NewBadgerDB(path string) (*BadgerDB, error) {
	//needed options to run in 32bit systems
	opts := badger.DefaultOptions(path).
		WithBaseTableSize(1 << 15).
		WithValueLogFileSize(1 << 20).
		WithSyncWrites(false)
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &BadgerDB{path, db}, nil
}

func (b *BadgerDB) Get(word []byte) ([]byte, error) {
	var result *[]byte
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(word)
		if err != nil {
			return err
		}
		err = item.Value(func(v []byte) error {
			result = &v
			return nil
		})
		return err
	})
	return *result, err
}

func (b *BadgerDB) Insert(key, value []byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(key, value)
		err := txn.SetEntry(e)
		return err
	})
}

func (b *BadgerDB) Delete(key []byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		return err
	})
}

func (b *BadgerDB) GetAll() ([]Entry, error) {
	var result []Entry
	err := b.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				result = append(result, Entry{k, v})
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return result, err
}

func (b *BadgerDB) Backup(bckFile io.Writer) error {
	_, err := b.db.Backup(bckFile, 0)
	return err
}

func (b *BadgerDB) Close() {
	b.db.Close()
}
