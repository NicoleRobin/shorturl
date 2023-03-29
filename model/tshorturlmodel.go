package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TShorturlModel = (*customTShorturlModel)(nil)

type (
	// TShorturlModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTShorturlModel.
	TShorturlModel interface {
		tShorturlModel
	}

	customTShorturlModel struct {
		*defaultTShorturlModel
	}
)

// NewTShorturlModel returns a model for the database table.
func NewTShorturlModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TShorturlModel {
	return &customTShorturlModel{
		defaultTShorturlModel: newTShorturlModel(conn, c, opts...),
	}
}
