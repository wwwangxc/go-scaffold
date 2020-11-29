package sqlite

import (
	"database/sql"
	"errors"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	store *dbStore

	ErrStoreNotFound = errors.New("[sqlite3] invalid store name.")
)

func init() {
	store = &dbStore{
		pool: make(map[string]*DB),
	}
}

// Append ..
func Append(storeName string, db *DB) (*DB, error) {
	if len(storeName) == 0 {
		return nil, ErrStoreNotFound
	}
	store.append(storeName, db)
	return db, nil
}

// Store ..
func Store(storeName string) *DB {
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
	pool map[string]*DB
	rw   sync.RWMutex
}

func (t *dbStore) append(storeName string, db *DB) {
	t.rw.Lock()
	defer t.rw.Unlock()
	db.StoreName = storeName
	t.pool[storeName] = db
}

func (t *dbStore) get(storeName string) (*DB, error) {
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
	if db.closed {
		return nil
	}
	t.rw.Lock()
	defer t.rw.Unlock()
	if db.closed {
		return nil
	}
	return db.Close()
}

type DB struct {
	*sql.DB

	StoreName string
	closed    bool
}

func newDB(conf *Config) *DB {
	db, err := sql.Open("sqlite3", conf.DSN)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return &DB{
		DB:     db,
		closed: false,
	}
}

// Close ..
func (t *DB) Close() error {
	if t.closed {
		return nil
	}
	if err := t.DB.Close(); err != nil {
		return err
	}
	t.closed = true
	return nil
}
