package xredis

import (
	"errors"
	"sync"
)

var (
	store *dbStore

	ErrStoreNotFound = errors.New("[xredis] invalid store name.")
)

func init() {
	store = &dbStore{
		pool: make(map[string]*Client),
	}
}

// Store ..
func Store(storeName string) *Client {
	db, err := store.get(storeName)
	if err != nil {
		panic(err)
	}
	return db
}

// Exist ..
func Exist(storeName string) bool {
	_, err := store.get(storeName)
	return err == nil
}

// Close ..
func Close(storeName string) error {
	return store.close(storeName)
}

// CloseAll ..
func CloseAll() {
	store.rw.Lock()
	defer store.rw.Unlock()
	for _, db := range store.pool {
		db.Close()
	}
}

type dbStore struct {
	pool map[string]*Client
	rw   sync.RWMutex
}

func (t *dbStore) append(storeName string, db *Client) {
	t.rw.Lock()
	defer t.rw.Unlock()
	t.pool[storeName] = db
}

func (t *dbStore) get(storeName string) (*Client, error) {
	if len(storeName) == 0 {
		return nil, ErrStoreNotFound
	}
	t.rw.RLock()
	defer t.rw.RUnlock()
	v, ok := t.pool[storeName]
	if !ok {
		return nil, ErrStoreNotFound
	}
	return v, nil
}

func (t *dbStore) close(storeName string) error {
	db, err := t.get(storeName)
	if err != nil {
		return err
	}
	if db.Closed {
		return nil
	}
	t.rw.Lock()
	defer t.rw.Lock()
	if db.Closed {
		return nil
	}
	return db.Close()
}
