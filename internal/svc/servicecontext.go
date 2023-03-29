package svc

import (
	"github.com/nicolerobin/shorturl/internal/config"
	"github.com/nicolerobin/shorturl/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	TShorturlModel model.TShorturlModel
	Redis          *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		TShorturlModel: model.NewTShorturlModel(conn, c.CacheConf),
		Redis:          redis.MustNewRedis(c.RedisConf),
	}
}
