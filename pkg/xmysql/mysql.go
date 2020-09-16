package xmysql

import (
	"errors"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	store *dbStore

	ErrStoreNotFound = errors.New("[xmysql] invalid store name.")
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
	if db.Closed {
		return nil
	}
	t.rw.Lock()
	defer t.rw.Unlock()
	if db.Closed {
		return nil
	}
	return db.Close()
}

type DB struct {
	*sqlx.DB

	StoreName string
	Closed    bool
}

func newDB(conf *Config) *DB {
	db, err := sqlx.Connect("mysql", conf.DSN)
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetConnMaxLifetime(conf.ConnMaxLifetime)
	return &DB{
		DB:     db,
		Closed: false,
	}
}

// Close ..
func (t *DB) Close() error {
	if t.Closed {
		return nil
	}
	if err := t.DB.Close(); err != nil {
		return err
	}
	t.Closed = true
	return nil
}
