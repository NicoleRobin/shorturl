package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	BaseUrl string
	Mysql   struct {
		DataSource string
	}
	CacheConf cache.CacheConf
	RedisConf redis.RedisConf
}
