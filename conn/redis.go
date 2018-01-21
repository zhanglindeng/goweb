package conn

import (
	"log"

	"github.com/gin-contrib/cache/persistence"
	"github.com/zhanglindeng/goweb/config"
)

var redisStore *persistence.RedisStore

func GetRedisConn() (*persistence.RedisStore, error) {

	if redisStore != nil {
		return redisStore, nil
	}

	log.Println("redisStore is nil")

	redisStore = persistence.NewRedisCache(config.RedisHost+":"+config.RedisPort, config.RedisPassword,
		persistence.DEFAULT)

	// test connection
	if err := redisStore.Set("PING", "PONG", persistence.DEFAULT); err != nil {
		return nil, err
	}

	return redisStore, nil
}
