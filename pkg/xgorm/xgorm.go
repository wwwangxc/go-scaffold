package xgorm

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
)

var (
	store *dbStore

	ErrStoreNotFound = errors.New("[xgorm] invalid store name.")
)

func init() {
	store = &dbStore{
		pool: make(map[string]*DB),
	}
}

// Store ..
func Store(storeName string) *DB {
	db, err := store.get(storeName)
	if err != nil {
		panic(err)
	}
	return db
}

// Close ..
func Close(storeName string) error {
	return store.close(storeName)
}

// Exist ..
func Exist(storeName string) bool {
	_, err := store.get(storeName)
	return err == nil
}

type dbStore struct {
	pool map[string]*DB
	rw   sync.RWMutex
}

func (t *dbStore) append(storeName string, db *DB) {
	t.rw.Lock()
	defer t.rw.Unlock()
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
	t.rw.Lock()
	defer t.rw.Lock()
	if db.Closed {
		return nil
	}
	return db.Close()
}

// DB ..
type DB struct {
	*gorm.DB

	Closed bool
}

func newDB(conf *Config) *DB {
	client, err := gorm.Open("mysql", conf.DSN)
	if err != nil {
		panic(err.Error())
	}
	client.DB().SetMaxIdleConns(conf.MaxIdleConns)
	client.DB().SetMaxOpenConns(conf.MaxOpenConns)
	client.DB().SetConnMaxLifetime(conf.ConnMaxLifetime)
	if err = client.DB().Ping(); err != nil {
		panic(err.Error())
	}
	return &DB{
		Closed: false,
		DB:     client,
	}
}

// Close ..
func (t *DB) Close() error {
	if t.Closed {
		return nil
	}
	t.Lock()
	defer t.Unlock()
	if err := t.DB.Close(); err != nil {
		return err
	}
	t.Closed = true
	return nil
}
