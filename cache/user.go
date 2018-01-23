package cache

import (
	"github.com/gin-contrib/cache/persistence"
	"github.com/zhanglindeng/goweb/util"
)

var userCachePrefix = "user_"

// SetUserTokenHash 缓存用户最新 token 的 md5 值
func SetUserTokenHash(email, hash string) error {
	c, err := getCacheConn()
	if err != nil {
		return err
	}

	return c.Set(keyUserTokenHash(email), hash, persistence.FOREVER)
}

func GetUserTokenHash(email string) (string, error) {
	c, err := getCacheConn()
	if err != nil {
		return "", err
	}
	var s string
	if err := c.Get(keyUserTokenHash(email), &s); err != nil {
		return "", err
	}

	return s, nil
}

func keyUserTokenHash(email string) string {
	return addKeyPrefix(userCachePrefix + util.Md5("keyUserTokenHash"+email))
}

func SetUserEmail2Id(email string, uid uint) error {
	c, err := getCacheConn()
	if err != nil {
		return err
	}

	return c.Set(keyUserEmail2Id(email), uid, persistence.FOREVER)
}

func GetUserEmail2Id(email string) (uint, error) {
	c, err := getCacheConn()
	if err != nil {
		return 0, err
	}
	var uid uint
	if err := c.Get(keyUserEmail2Id(email), &uid); err != nil {
		return 0, err
	}

	return uid, nil
}

func keyUserEmail2Id(email string) string {
	return addKeyPrefix(userCachePrefix + util.Md5("keyUserEmail2Id"+email))
}
