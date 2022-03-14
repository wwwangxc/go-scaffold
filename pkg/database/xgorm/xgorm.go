package xgorm

import (
	"sync"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	store *dbStore

	ErrStoreNotFound = errors.New("[xgorm] invalid store name.")
)

func init() {
	store = &dbStore{
		pool: make(map[string]*gorm.DB),
	}
}

// Store ..
func Store(storeName string) *gorm.DB {
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

type dbStore struct {
	pool map[string]*gorm.DB
	rw   sync.RWMutex
}

func (t *dbStore) append(storeName string, db *gorm.DB) {
	t.rw.Lock()
	defer t.rw.Unlock()
	t.pool[storeName] = db
}

func (t *dbStore) get(storeName string) (*gorm.DB, error) {
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

func newDB(conf *Config) *gorm.DB {
	c, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db, err := c.DB()
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetConnMaxLifetime(conf.ConnMaxLifetime)

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	return c
}
