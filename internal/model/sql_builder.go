package model

import (
	"context"
	"fmt"
	xgorm "go-scaffold/pkg/database/xgorm"

	"gorm.io/gorm"
)

type sqlBuilder struct {
	store string

	whereConditions   []string
	whereArgs         []interface{}
	orderByConditions []string
	limit             int
}

func newSQLBuilder(store string, opts ...Option) (*sqlBuilder, error) {
	builder := &sqlBuilder{
		store: store,
	}

	for _, opt := range opts {
		if err := opt(builder); err != nil {
			return nil, err
		}
	}

	return builder, nil
}

func (s *sqlBuilder) mustBuild(ctx context.Context) *gorm.DB {
	db, _ := s.build(ctx)
	return db
}

func (s *sqlBuilder) build(ctx context.Context) (*gorm.DB, error) {
	if s == nil {
		return nil, fmt.Errorf("invalid sql builder")
	}

	return s.buildWithDBClient(ctx, xgorm.Store(s.store))
}

func (s *sqlBuilder) mustBuildWithDBClient(ctx context.Context, db *gorm.DB) *gorm.DB {
	db, _ = s.buildWithDBClient(ctx, db)
	return db
}

func (s *sqlBuilder) buildWithDBClient(ctx context.Context, db *gorm.DB) (*gorm.DB, error) {
	if s == nil {
		return nil, fmt.Errorf("invalid sql builder")
	}

	if db == nil {
		return nil, fmt.Errorf("invalid db client")
	}

	db = db.WithContext(ctx)

	db, err := s.buildWhere(db)
	if err != nil {
		return nil, fmt.Errorf("build where condition fail. err:%w", err)
	}

	db, err = s.buildOrderBy(db)
	if err != nil {
		return nil, fmt.Errorf("build order by condition fail. err:%w", err)
	}

	db, err = s.buildLimit(db)
	if err != nil {
		return nil, fmt.Errorf("build limit condition fail. err:%w", err)
	}

	return db, nil
}

func (s *sqlBuilder) buildWhere(db *gorm.DB) (*gorm.DB, error) {
	if s == nil {
		return nil, fmt.Errorf("invalid sql builder")
	}

	if db == nil {
		return nil, fmt.Errorf("invalid db client")
	}

	if len(s.whereConditions) == 0 {
		return db, nil
	}

	if len(s.whereArgs) != len(s.whereConditions) {
		return nil, fmt.Errorf("invalid where args. expected:%d actual:%d", len(s.whereConditions), len(s.whereArgs))
	}

	for i, v := range s.whereConditions {
		db.Where(v, s.whereArgs[i])
	}

	return db, nil
}

func (s *sqlBuilder) buildOrderBy(db *gorm.DB) (*gorm.DB, error) {
	if s == nil {
		return nil, fmt.Errorf("invalid sql builder")
	}

	if db == nil {
		return nil, fmt.Errorf("invalid db client")
	}

	if len(s.orderByConditions) == 0 {
		return db, nil
	}

	for _, v := range s.orderByConditions {
		db.Order(v)
	}

	return db, nil
}

func (s *sqlBuilder) buildLimit(db *gorm.DB) (*gorm.DB, error) {
	if s == nil {
		return nil, fmt.Errorf("invalid sql builder")
	}

	if db == nil {
		return nil, fmt.Errorf("invalid db client")
	}

	if s.limit < 0 {
		return nil, fmt.Errorf("invalid limit. expected greater or equal than 0 actual:%d", s.limit)
	}

	if s.limit == 0 {
		return db, nil
	}

	db.Limit(s.limit)

	return db, nil
}

func (s *sqlBuilder) appendWhereCondition(condition string, arg interface{}) error {
	if s == nil {
		return fmt.Errorf("invalid sql builder")
	}

	if condition == "" {
		return fmt.Errorf("invalid condition")
	}

	if arg == nil {
		return fmt.Errorf("invalid arg")
	}

	s.whereConditions = append(s.whereConditions, condition)
	s.whereArgs = append(s.whereArgs, arg)

	return nil
}

func (s *sqlBuilder) appendOrderByCondition(condition string) error {
	if s == nil {
		return fmt.Errorf("invalid sql builder")
	}

	if condition == "" {
		return fmt.Errorf("invalid condition")
	}

	s.orderByConditions = append(s.orderByConditions, condition)

	return nil
}

func (s *sqlBuilder) setLimit(limit int) error {
	if s == nil {
		return fmt.Errorf("invalid sql builder")
	}

	if s.limit < 1 {
		return fmt.Errorf("invalid limit. expected greater than 0 actual:%d", limit)
	}

	s.limit = limit

	return nil
}
