package cache

import (
	"github.com/gin-contrib/cache/persistence"
	"github.com/zhanglindeng/goweb/config"
	"github.com/zhanglindeng/goweb/conn"
)

func getCacheConn() (*persistence.RedisStore, error) {
	return conn.GetRedisConn()
}

func addKeyPrefix(s string) string {
	return config.CachePrefix + s
}
