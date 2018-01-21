package config

import (
	"os"
	"strconv"

	redis2 "github.com/garyburd/redigo/redis"
)

var (
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDatabase int
)

func redis() error {

	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	redisDatabaseStr := os.Getenv("REDIS_DATABASE")

	redisDatabase, err := strconv.Atoi(redisDatabaseStr)
	if err != nil {
		return err
	}

	RedisDatabase = redisDatabase

	// test conn
	c, err := redis2.Dial("tcp", RedisHost+":"+RedisPort, redis2.DialPassword(RedisPassword),
		redis2.DialDatabase(RedisDatabase))
	if err != nil {
		return err
	}
	defer c.Close()

	return nil
}
