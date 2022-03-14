package model

// Option ...
type Option func(*sqlBuilder) error

func WithIDEqual(id int) Option {
	return newWhereOption("id=?", id)
}

// WithUserNameEuqa WHERE user_name=?
func WithUserNameEuqa(userName string) Option {
	return newWhereOption("user_name", userName)
}

// WithOrderByCreateAt ORDER BY create_at
func WithOrderByCreateAt() Option {
	return newOrderByOption("create_at")
}

// WithOrderByCreateAtDESC ORDER BY create_at DESC
func WithOrderByCreateAtDESC() Option {
	return newOrderByOption("create_at DESC")
}

// WithOrderByUpdateAt ORDER BY update_at
func WithOrderByUpdateAt() Option {
	return newOrderByOption("update_at")
}

// WithOrderByUpdateAtDESC ORDER BY update_at DESC
func WithOrderByUpdateAtDESC() Option {
	return newOrderByOption("update_at DESC")
}

func newWhereOption(condition string, arg interface{}) Option {
	return func(s *sqlBuilder) error {
		return s.appendWhereCondition(condition, arg)
	}
}

func newOrderByOption(condition string) Option {
	return func(s *sqlBuilder) error {
		return s.appendOrderByCondition(condition)
	}
}
